import BasicLayout from '@/layouts/BasicLayout';

export default [
  {
    path: '/article',
    redirect: '/article/list',
    component: BasicLayout,
    name: 'Article',
    meta: { title: '文章', icon: 'view-module' },
    children: [
      {
        path: 'list',
        name: 'ArticleList',
        component: () => import('@/views/article/ArticleList.vue'),
        meta: { title: '所有文章', icon: 'view-module' },
      },
      {
        path: 'edit',
        name: 'ArticleEdit',
        component: () => import('@/views/article/ArticleEdit.vue'),
        meta: { title: '写文章', icon: 'edit-1' },
      },
    ],
  },
]