import BasicLayout from '@/layouts/BasicLayout';

export default [
  {
    path: '/comment',
    component: BasicLayout,
    meta: { title: '评论', icon: 'edit-1', single: true },
    children: [
      {
        path: '',
        name: 'Comment',
        component: () => import('@/views/comment/CommentList.vue'),
      },
    ],
  },
];
