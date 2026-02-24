import VueRouter from 'vue-router';
// 静态路由
import staticRoutes from '@/router/static';
// 侧边栏路由
import dashBoardRoutes from '@/router/modules/dashboard';
import articleRoutes from '@/router/modules/article';
import linkRoutes from '@/router/modules/links';
import appearanceRoutes from '@/router/modules/appearance';
import commentRoutes from '@/router/modules/comment';
import journalRoutes from '@/router/modules/journal';
import settingsRoutes from '@/router/modules/settings';

// 如果需要新增只有一级路由的侧边栏请参考dashboard路由的写法：./modules/dashboard.js
// 显示在侧边栏中的路由，用于生成侧边栏
export const menuRoutes = [
  ...dashBoardRoutes,
  ...articleRoutes,
  ...linkRoutes,
  ...journalRoutes,
  ...commentRoutes,
  ...appearanceRoutes,
  ...settingsRoutes,
]

// 合并所有路由
export const allRoutes = [
  ...staticRoutes,
  ...menuRoutes,
];

export default new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: allRoutes,
});
