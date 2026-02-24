import BasicLayout from '@/layouts/BasicLayout';

export default [
  {
    path: '/settings',
    component: BasicLayout,
    meta: { title: '系统', icon: 'setting' },
    children: [
      {
        path: 'settings',
        name: 'Settings',
        component: () => import('@/views/system/Settings.vue'),
        meta: { title: '博客设置', icon: 'setting' },
      },
      {
        path: 'userinfo',
        name: 'UserInfo',
        component: () => import('@/views/system/UserInfo.vue'),
        meta: { title: '个人信息', icon: 'user' },
      },
    ],
  },
];
