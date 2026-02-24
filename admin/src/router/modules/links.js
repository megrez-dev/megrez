import BasicLayout from '@/layouts/BasicLayout';

export default [
  {
    path: '/links',
    component: BasicLayout,
    meta: { title: '友链', icon: 'link', single: true },
    children: [
      {
        path: '',
        name: 'Links',
        component: () => import('@/views/links/Links.vue'),
      },
    ],
  },
]
