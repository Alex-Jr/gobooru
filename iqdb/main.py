from http.server import BaseHTTPRequestHandler, HTTPServer
import json
import image_hash
import numpy as np
import time
import os
import cgi

from sqlalchemy import create_engine, text, Column, Integer, LargeBinary, ARRAY, Text
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

engine = create_engine(
  'postgresql://{DB_USER}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_DATABASE}?sslmode={DB_SSLMODE}'.format(
    DB_USER=os.environ.get('DB_USER', 'user'),
    DB_PASSWORD=os.environ.get('DB_PASSWORD', 'password'),
    DB_HOST=os.environ.get('DB_HOST', 'localhost'),
    DB_PORT=os.environ.get('DB_PORT', '5450'),
    DB_DATABASE=os.environ.get('DB_DATABASE', 'database'),
    DB_SSLMODE=os.environ.get('DB_SSLMODE', 'disable')
  ), 
  pool_timeout=60
)

Base = declarative_base()
class PostSignature(Base):
    __tablename__ = 'post_signatures'

    post_id = Column(Integer, primary_key=True)
    signature = Column(LargeBinary)
    words = Column(ARRAY(Text))

Session = sessionmaker(bind=engine)

class MyHandler(BaseHTTPRequestHandler):
  def do_GET(self):
    self.send_response(200)
    self.send_header('Content-type', 'application/json')
    self.end_headers()
    self.wfile.write(json.dumps({'hello': 'world', 'received': 'ok'}).encode('utf-8'))

  def do_POST(self):
    if self.path == '/add':
      self.handle_add()
    elif self.path == '/search':
      self.handle_search()
    elif self.path == '/search-file':
      self.handle_search_file()
    else:
      self.send_response(404)
      self.end_headers()

  def handle_add(self):
    start = time.time()

    content_length = int(self.headers['Content-Length'])
    post_data = self.rfile.read(content_length)
    data = json.loads(post_data.decode('utf-8'))
    try: 
      session = Session()

      with open(data['filePath'], "rb") as f:
        postId = data['postId']
        signature = image_hash.generate_signature(f.read())
        words = image_hash.generate_words(signature)

        
        new_post = PostSignature(post_id=postId, signature=image_hash.pack_signature(signature), words=words)

        session.add(new_post)

        session.commit()

        query = text("""
          SELECT 
            ps."post_id", 
            ps."signature", 
            count(a."query") AS "score"
          FROM 
            "post_signatures" AS ps, 
            unnest(ps."words", :q) AS a("word", "query")
          WHERE 
            a."word" = a."query"::text
            and ps."post_id" <> :postId
          GROUP BY 
            ps."post_id"
          ORDER BY 
            score DESC 
          LIMIT 100;
        """)

        candidates = session.execute(query, {"q": words, "postId": postId})

        data = tuple(
          zip(
            *[
              (post_id, image_hash.unpack_signature(packedsig))
              for post_id, packedsig, score in candidates
            ]
          )
        )
        response = {}

        if data:
          candidate_post_ids, sigarray = data
          distances = image_hash.normalized_distance(sigarray, signature)

          json_array = []

          i = 0
          for d in distances:
            json_array.append({
              "postId": candidate_post_ids[i],
              "similarity": (1 - d) * 100
            })

            i+=1
    
          response = {
            "similarities": json_array,
          }
      
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()
        self.wfile.write(json.dumps(response).encode('utf-8'))
    except Exception as e:
      session.rollback()
      print(e)
      self.send_response(500)
      self.send_header('Content-type', 'application/json')
      self.end_headers()
      self.wfile.write(json.dumps({"err": str(e)}).encode('utf-8'))

    finally:
      end = time.time()
      print('Elapsed: ' + str(end - start))
      session.close()

  # def handle_search(self):
  #   content_length = int(self.headers['Content-Length'])
  #   post_data = self.rfile.read(content_length)
  #   data = json.loads(post_data.decode('utf-8'))

  #   try: 
  #     session = Session()

  #     with open(data['filePath'], "rb") as f:
  #       signature = image_hash.generate_signature(f.read())
  #       words = image_hash.generate_words(signature)

  #       query = text("""
  #         SELECT 
  #           ps."postId", 
  #           ps."signature", 
  #           count(a."query") AS "score"
  #         FROM 
  #           "PostSignature" AS ps, 
  #           unnest(ps."words", :q) AS a("word", "query")
  #         WHERE 
  #           a."word" = a."query"::text
  #         GROUP BY 
  #           ps."postId"
  #         ORDER BY 
  #           score DESC 
  #         LIMIT 100;
  #       """)

  #       candidates = session.execute(query, {"q": words})

  #       data = tuple(
  #         zip(
  #           *[
  #             (post_id, image_hash.unpack_signature(packedsig))
  #             for post_id, packedsig, score in candidates
  #           ]
  #         )
  #       )
  #       response = {}

  #       if data:
  #         candidate_post_ids, sigarray = data
  #         distances = image_hash.normalized_distance(sigarray, signature)

  #         json_array = []

  #         i = 0
  #         for d in distances:
  #           json_array.append({
  #             "postId": candidate_post_ids[i],
  #             "similarity": (1 - d) * 100
  #           })

  #           i+=1
    
  #         response = {
  #           "similarities": json_array,
  #         }
          
  #       self.send_response(200)
  #       self.send_header('Content-type', 'application/json')
  #       self.end_headers()
  #       self.wfile.write(json.dumps(response).encode('utf-8'))
  #   except Exception as e:
  #     print(e)
  #     session.rollback()
  #   finally:
  #     session.close()

  # def handle_search_file(self):
    # form = cgi.FieldStorage(
    #   fp=self.rfile,
    #   headers=self.headers,
    #   environ={'REQUEST_METHOD': 'POST'}
    # )

    # for field in form.keys():
    #   field_item = form[field]
    #   if field_item.filename:
    #     file_data = field_item.file.read()
        
    #     print(f"Received file: {field_item.filename}")

    #   try: 
    #     session = Session()

    #     signature = image_hash.generate_signature(file_data)
    #     words = image_hash.generate_words(signature)

    #     query = text("""
    #       SELECT 
    #         ps."postId", 
    #         ps."signature", 
    #         count(a."query") AS "score"
    #       FROM 
    #         "PostSignature" AS ps, 
    #         unnest(ps."words", :q) AS a("word", "query")
    #       WHERE 
    #         a."word" = a."query"::text
    #       GROUP BY 
    #         ps."postId"
    #       ORDER BY 
    #         score DESC 
    #       LIMIT 100;
    #     """)

    #     candidates = session.execute(query, {"q": words})

    #     data = tuple(
    #       zip(
    #         *[
    #           (post_id, image_hash.unpack_signature(packedsig))
    #           for post_id, packedsig, score in candidates
    #         ]
    #       )
    #     )
    #     response = {}

    #     if data:
    #       candidate_post_ids, sigarray = data
    #       distances = image_hash.normalized_distance(sigarray, signature)

    #       json_array = []

    #       i = 0
    #       for d in distances:
    #         json_array.append({
    #           "postId": candidate_post_ids[i],
    #           "similarity": (1 - d) * 100
    #         })

    #         i+=1
    
    #       response = {
    #         "similarities": json_array,
    #       }
          
    #     self.send_response(200)
    #     self.send_header('Content-type', 'application/json')
    #     self.end_headers()
    #     self.wfile.write(json.dumps(response).encode('utf-8'))
    #   except Exception as e:
    #     print(e)
    #     session.rollback()
    #   finally:
    #     session.close()
def run(server_class=HTTPServer, handler_class=MyHandler, port=8000):
  server_address = ('', port)
  httpd = server_class(server_address, handler_class)
  print(f"Starting server on port {port}...")
  httpd.serve_forever()


run()
