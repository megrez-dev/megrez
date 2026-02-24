// 不在侧边栏中显示的路由
export default [
  {
    path: '/',
    name: 'Index',
    redirect: '/dashboard',
  },
  {
    path: '/w+',
    redirect: '/404',
  },
  {
    path: '/login',
    name: 'Login',
    meta: { title: '登录' },
    component: () => import('@/views/static/Login'),
  },
  {
    path: '/install',
    name: 'Install',
    meta: { title: '安装向导' },
    component: () => import('@/views/static/Install'),
  },
]