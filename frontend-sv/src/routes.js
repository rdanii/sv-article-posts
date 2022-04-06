import React from 'react'

const Dashboard = React.lazy(() => import('./views/dashboard/Dashboard'))

const Publish = React.lazy(() => import('./views/pages/articles/Publish'))
const Draft = React.lazy(() => import('./views/pages/articles/Draft'))
const Thrash = React.lazy(() => import('./views/pages/articles/Thrash'))
const AddArticle = React.lazy(() => import('./views/pages/articles/add-article/AddArticle'))
const PreviewArticle = React.lazy(() => import('./views/pages/articles/get-article/PreviewArticle'))
const EditArticle = React.lazy(() => import('./views/pages/articles/edit-article/EditArticle'))

const routes = [
  { path: '/', exact: true, name: 'Home' },
  { path: '/dashboard', name: 'Dashboard', element: Dashboard },
  { path: '/all-post', name: 'Publish', element: Publish },
  { path: '/thrash', name: 'Thrash', element: Thrash },
  { path: '/draft', name: 'Draft', element: Draft },
  { path: '/add-new', name: 'Add New', element: AddArticle },
  { path: '/preview', name: 'Preview', element: PreviewArticle },
  { path: '/edit/:id', name: 'Edit', element: EditArticle },
]

export default routes
