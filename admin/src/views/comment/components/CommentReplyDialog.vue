<template>
  <t-dialog
    header="回复评论"
    :visible="visible"
    confirmBtn="回复"
    :cancelBtn="null"
    :onClose="close"
    :onConfirm="replyComment"
  >
    <t-textarea
      v-model="newComment.content"
      placeholder="暂不支持 Markdown 评论"
      :autosize="{ minRows: 10 }"
    ></t-textarea>
  </t-dialog>
</template>

<script>
export default {
  name: "CommentReplyDialog",
  data() {
    return {
      visible: false,
      newComment: {
        content: "",
      },
    };
  },
  props: {
    comment: {
      required: true,
    },
  },
  methods: {
    open() {
      this.visible = true;
    },
    close() {
      this.visible = false;
    },
    replyComment() {
      if (this.newComment.content.length === 0) {
        this.$message.warning("评论内容不能为空");
        return;
      }
      if (this.comment.rootID === 0) {
        this.newComment.rootID = this.comment.id;
      } else {
        this.newComment.rootID = this.comment.rootID;
      }
      this.newComment.parentID = this.comment.id;
      this.newComment.articleID = this.comment.article.id;
      this.newComment.pageID = this.comment.page.id;
      this.newComment.type = this.comment.type;
      this.$request
        .post("/comment", this.newComment)
        .then(() => {
          this.$message.success("回复成功");
          this.$emit("replySuccess");
          this.close();
        })
        .catch(() => {
          this.$message.error("回复失败");
        });
    },
  },
};
</script>
<style>
</style>