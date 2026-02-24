<template>
  <PageView>
    <template slot="content">
      <div class="article-list-container">
        <div class="card-container">
          <t-table
            :data="commentList"
            :columns="columns"
            rowKey="id"
            verticalAlign="middle"
            :loading="isCommentListLoading"
            :pagination="pagination"
            @change="rehandleChange"
          >
            <template #title="{ row }">
              <a
                class="t-button-link"
                :href="'/article/' + row.article.id"
                target="_blank"
                v-if="row.type === COMMENT_TYPE.ARTICLE"
                >{{ row.article.title }}</a
              >
              <a
                class="t-button-link"
                :href="'/' + row.page.slug"
                target="_blank"
                v-else
                >{{ row.page.name }}</a
              >
            </template>
            <template #author="{ row }">
              <a
                class="t-button-link"
                :href="row.url"
                target="_blank"
                v-if="row.url != ''"
                >{{ row.author }}</a
              >
            </template>
            <template #status="{ row }">
              <span v-if="row.status === 0">
                <t-badge dot :offset="[8, -5]" color="green"> </t-badge
                >已发布</span
              >
              <span v-else-if="row.status === 1"
                ><t-badge dot :offset="[8, -5]" color="orange"> </t-badge
                >待审核</span
              >
              <span v-else>
                <t-badge dot :offset="[8, -5]" color="red"> </t-badge>未知</span
              >
            </template>
            <template #op="slotProps">
              <a class="t-button-link" @click="handleClickReply(slotProps)"
                >回复</a
              >
              <t-divider layout="vertical" />
              <a class="t-button-link" @click="handleClickDelete(slotProps)"
                >删除</a
              >
            </template>
          </t-table>
        </div>
        <CommentReplyDialog
          ref="commentReplyDialog"
          :comment="replyComment"
          @replySuccess="onReplySuccess"
        ></CommentReplyDialog>
      </div>
    </template>
  </PageView>
</template>

<script>
import { timeAgo } from "@/utils/datetime.js";
import PageView from "@/layouts/PageView";
import CommentReplyDialog from "./components/CommentReplyDialog.vue";

export default {
  name: "CommentList",
  data() {
    return {
      COMMENT_TYPE: {
        ARTICLE: "article",
        PAGE: "page",
        // JOURNAL: "journal",
      },
      commentList: [],
      isCommentListLoading: false,
      columns: [
        {
          colKey: "title",
          title: "文章/页面",
        },
        {
          colKey: "content",
          title: "内容",
        },
        {
          colKey: "author",
          title: "作者",
        },
        {
          colKey: "mail",
          title: "邮箱",
        },
        {
          colKey: "ip",
          title: "IP",
          width: "150px",
        },
        {
          colKey: "status",
          title: "状态",
          width: "100px",
        },
        {
          colKey: "createTime",
          title: "发布时间",
          cell(h, { row: { createTime } }) {
            return timeAgo(createTime);
          },
        },
        {
          fixed: "right",
          colKey: "op",
          title: "操作",
        },
      ],
      replyComment: null,
      pagination: {
        current: 1,
        pageSize: 10,
      },
    };
  },
  mounted() {
    this.fetchComments();
  },
  methods: {
    fetchComments(pagination = this.pagination) {
      this.isCommentListLoading = true;
      const { current, pageSize } = pagination;
      this.$request
        .get("comments?pageNum=" + current + "&pageSize=" + pageSize)
        .then((res) => {
          this.commentList = res.data.list ? res.data.list : [];
          this.pagination = {
            ...pagination,
            total: res.data.total,
          };
        })
        .catch(() => {
          this.$message.error("获取评论列表失败");
        })
        .finally(() => {
          this.isCommentListLoading = false;
        });
    },
    rehandleChange(pageInfo) {
      this.pagination.current = pageInfo.pagination.current;
      this.pagination.pageSize = pageInfo.pagination.pageSize;
      this.fetchComments();
    },
    handleClickReply(slotProps) {
      this.replyComment = slotProps.row;
      this.$refs.commentReplyDialog.open();
    },
    handleClickDelete(slotProps) {
      this.$request
        .delete("comment/" + slotProps.row.id)
        .then(() => {
          this.$message.success("删除成功");
          this.fetchComments();
        })
        .catch(() => {
          this.$message.warning("删除失败");
          // TODO: 换成修改按钮状态
        });
    },
    onReplySuccess() {
      this.fetchComments();
    },
  },
  components: { PageView, CommentReplyDialog },
};
</script>
<style lang="less" scoped>
.article-list-container {
  .card-container {
    padding: 16px 24px;
    background-color: @bg-color-container;
    border-radius: 2px;
    width: 100%;
    display: flex;
    flex-direction: column;
    .left-operation-container {
      padding: 0 0 6px 0;
      margin-bottom: 16px;
    }
  }
}
</style>