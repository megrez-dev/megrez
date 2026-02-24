import BasicLayout from '@/layouts/BasicLayout';

// 由于在业务组件之上需要指明布局组件
// 所以必须要有一个根路径('/dashboard')用于指明当前页面使用的布局组件
// children只用于引入业务组件，即页面中间部分
export default [
  {
    path: '/dashboard',
    component: BasicLayout,
    meta: { title: '仪表盘', icon: 'dashboard', single: true },
    children: [
      {
        path: '',
        // name必须在此指定，否则vue将报警告：
        // Named Route 'Dashboard' has a default child route. When navigating to this named route (:to="{name: 'Dashboard'"), the default child route will not be rendered. Remove the name from this route and use the name of the default child route for named links instead.
        name: 'Dashboard',
        component: () => import('@/views/dashboard/Dashboard.vue'),
      },
    ],
  },
];
