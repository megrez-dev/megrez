export default {
    state: {
        collapsed: false,
        isDark: false,
        attachDetailDrawerVisible: false,
        attachUploadDialogVisible: false
    },
    mutations: {
        TOGGLE_COLLAPSED(state) {
            state.collapsed = !state.collapsed
        },
        TOGGLE_THEME(state) {
            state.isDark = !state.isDark
        },
        OPEN_ATTACH_DETAIL_DRAWER(state) {
            state.attachDetailDrawerVisible = true
        },
        CLOSE_ATTACH_DETAIL_DRAWER(state) {
            state.attachDetailDrawerVisible = false
        },
        OPEN_ATTACH_UPLOAD_DIALOG(state) {
            state.attachUploadDialogVisible = true
        },
        CLOSE_ATTACH_UPLOAD_DIALOG(state) {
            state.attachUploadDialogVisible = false
        }
    }
}