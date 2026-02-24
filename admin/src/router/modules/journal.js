import BasicLayout from '@/layouts/BasicLayout';

export default [
  {
    path: '/journal',
    component: BasicLayout,
    meta: { title: '日志', icon: 'edit-1', single: true },
    children: [
      {
        path: '',
        name: 'Journal',
        component: () => import('@/views/journal/JournalList.vue'),
    }
    ],
  },
];
