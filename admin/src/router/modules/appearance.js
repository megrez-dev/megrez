import BasicLayout from '@/layouts/BasicLayout';

export default [
  {
    path: '/appearance',
    redirect: '/appearance/theme',
    component: BasicLayout,
    name: 'Appearance',
    meta: { title: '外观', icon: 'view-module' },
    children: [
      {
        path: 'theme',
        name: 'Theme',
        component: () => import('@/views/appearance/Theme.vue'),
        meta: { title: '主题', icon: 'view-module' },
      },
      {
        path: 'setting',
        name: 'ThemeSetting',
        component: () => import('@/views/appearance/ThemeSetting.vue'),
        meta: { title: '主题设置', icon: 'edit-1' },
      },
      // {
      //     path: 'edit',
      //     component: () => import('@/views/appearance/ThemeEdit.vue'),
      //     name: 'ThemeEdit',
      //     meta: { title: '主题编辑', icon: 'edit-1' },
      // },
    ],
  },
];
