export const EDIT_MODE = {
  create: 0,
  edit: 1,
};
// TODO: 修改 placeholder 图片
export const DEFAULT_COVER_URL = 'http://119.91.26.252:8090/images/placeholder.jpg';

export const THEME_SETTING = {
  light: {
    theme: 'classic', // 编辑器主题
    contentTheme: 'light', // 内容主题
  },
  dark: {
    theme: 'dark',
    contentTheme: 'dark',
  }
}

export const INSERT_IMG_CONFIGS = {
  name: 'megrez-image',
  tip: '插入图片',
  tipPosition: 'nw',
  icon: '<svg t="1659202211997" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2263" width="200" height="200"><path d="M736 448c53 0 96-43 96-96 0-53-43-96-96-96-53 0-96 43-96 96C640 405 683 448 736 448z" p-id="2264"></path><path d="M904 128 120 128c-31.2 0-56 25.4-56 56.6l0 654.8c0 31.2 24.8 56.6 56 56.6l784 0c31.2 0 56-25.4 56-56.6L960 184.6C960 153.4 935.2 128 904 128zM697.8 523.4c-6-7-15.2-12.4-25.6-12.4-10.2 0-17.4 4.8-25.6 11.4l-37.4 31.6c-7.8 5.6-14 9.4-23 9.4-8.6 0-16.4-3.2-22-8.2-2-1.8-5.6-5.2-8.6-8.2L448 430.6c-8-9.2-20-15-33.4-15-13.4 0-25.8 6.6-33.6 15.6L128 736.4 128 215.4c2-13.6 12.6-23.4 26.2-23.4l715.4 0c13.8 0 25 10.2 25.8 24l0.6 520.8L697.8 523.4z" p-id="2265"></path></svg>',
};

export const VDITOR_TOOLBAR = [
  'outline',
  'undo',
  'redo',
  'headings',
  'bold',
  'italic',
  'strike',
  'link',
  // 此处插入megrez-image
  'list',
  'ordered-list',
  'table',
  'check',
  'outdent',
  'indent',
  'quote',
  'line',
  'code',
  'inline-code',
  'code-theme',
  'content-theme',
  'export',
  'fullscreen',
];

export const VDITOR_BASE_CONFIGS = {
  height: 'calc(100vh - 250px)',
  icon: 'material',
  mode: 'wysiwyg',
  outline: {
    enable: true,
  },
  cache: {
    enable: false,
  },
};

export const SKELETON_ROW_COL = [
  { width: '30%', },
  1, 1,
  { width: '30%' },
  1, 1, 1,
  { width: '80%' },
  1,
  { width: '60%' },
  1, 1, 1, 1,
  { width: '80%' },
  { width: '10%' },
]

export const ARTICLE_STATUS_MAP = {
  0: {
    color: 'green',
    text: '已发布',
    theme: 'primary',
  },
  1: {
    color: 'orange',
    text: '草稿',
    theme: 'warning',
  },
};