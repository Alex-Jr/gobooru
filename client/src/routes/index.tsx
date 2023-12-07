import { BrowserRouter, Route, Routes } from "react-router-dom";

import { HomePage } from "pages/home/home-page";
import { LoginPage } from "pages/login/login-page";
import { PoolPage } from "pages/pool/pool-page";
import { PoolListPage } from "pages/pool-list/pool-list-page";
import { PostPage } from "pages/post/post-page";
import { PostListPage } from "pages/post-list/post-list-page";
import { PostNewPage } from "pages/post-new/post-new-page";
import { PostNewPageBatch } from "pages/post-new-batch/post-new-page-batch";
import { TagPage } from "pages/tag/tag-page";
import { TagListPage } from "pages/tag-list/tag-list-page";
import { AuthenticatedRoute } from "shared/components/navigation/authenticated-route";
import { BaseRoute } from "shared/components/navigation/base-route";

export const MainRouter = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={
            <BaseRoute>
              <HomePage />
            </BaseRoute>
          }
        />

        <Route
          path="/login"
          element={
            <BaseRoute>
              <LoginPage />
            </BaseRoute>
          }
        />

        <Route
          path="/posts"
          element={
            <BaseRoute>
              <PostListPage />
            </BaseRoute>
          }
        />

        <Route
          path="/posts/new"
          element={
            <AuthenticatedRoute>
              <PostNewPage />
            </AuthenticatedRoute>
          }
        />

        <Route
          path="/posts/new/batch"
          element={
            <AuthenticatedRoute>
              <PostNewPageBatch />
            </AuthenticatedRoute>
          }
        />

        <Route
          path="/posts/:id"
          element={
            <BaseRoute>
              <PostPage />
            </BaseRoute>
          }
        />

        <Route
          path="/pools"
          element={
            <BaseRoute>
              <PoolListPage />
            </BaseRoute>
          }
        />

        <Route
          path="/pools/:id"
          element={
            <BaseRoute>
              <PoolPage />
            </BaseRoute>
          }
        />

        <Route
          path="/tags"
          element={
            <BaseRoute>
              <TagListPage />
            </BaseRoute>
          }
        />

        <Route
          path="/tags/:id"
          element={
            <BaseRoute>
              <TagPage />
            </BaseRoute>
          }
        />
      </Routes>
    </BrowserRouter>
  );
};
