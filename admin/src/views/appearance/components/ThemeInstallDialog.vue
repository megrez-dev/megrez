<template>
  <div>
    <t-dialog
      :visible="visible"
      :header="null"
      :confirmBtn="null"
      :cancelBtn="null"
      :onClose="close"
    >
      <div class="upload-container">
        <t-upload
          v-model="files"
          :onSuccess="handleInstallSuccess"
          :onFail="handleInstallFail"
          :requestMethod="installMethod"
          draggable
        />
      </div>
    </t-dialog>
  </div>
</template>
<script>
export default {
  name: "themeInstallDialog",
  data() {
    return {
      visible: false,
      files: [],
    };
  },
  methods: {
    open() {
      this.visible = true;
    },
    close() {
      this.visible = false;
    },
    handleInstallFail() {
      this.$message.error("安装失败");
    },
    handleInstallSuccess() {
      this.$message.success("安装成功");
      this.$emit("installSuccess");
      this.files = [];
      this.close();
    },
    installMethod(file) {
      return new Promise((resolve) => {
        // file.percent 用于控制上传进度，如果不希望显示上传进度，则不对 file.percent 设置值即可。
        // 如果代码规范不能设置 file.percent，也可以设置 this.files
        file.percent = 0;
        let data = new FormData();
        data.append("file", file.raw);
        this.$request
          .post("theme/install", data, {
            headers: { "Content-Type": "multipart/form-data" },
          })
          .then((res) => {
            if (res.status === 0) {
              resolve({
                status: "success",
                response: {},
              });
            }
          })
          .catch(() => {
            resolve({
              status: "fail",
              error: "上传失败，请检查文件是否符合规范",
            });
          });
      });
    },
  },
};
</script>

<style lang="less" scoped>
.upload-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
}
</style>