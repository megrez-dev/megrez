<template>
  <PageView>
    <template slot="content">
      <t-list
        :split="true"
        :async-loading="asyncLoading"
        @load-more="fetchJournals"
      >
        <template #header>
          <div class="left-operation-container">
            <t-button @click="onClickNew">
              <add-icon slot="icon" />
              写日志
            </t-button>
          </div>
        </template>
        <t-list-item v-for="(item, i) in journalList" :key="i">
          <template #content>
            <div class="journal-list-item-container">
              <div class="journal-list-item-content">
                <span>{{ item.content }}</span>
              </div>
              <div class="journal-list-item-images">
                <div
                  class="journal-list-image-box"
                  v-for="(image, j) in item.images"
                  :key="j"
                  :style="{ backgroundImage: `url(${image})` }"
                  @mouseenter="mouseEnter(i, j)"
                  @mouseleave="mouseLeave(i, j)"
                >
                  <div
                    class="journal-list-image-mask"
                    v-show="maskVisible[i][j]"
                    @click="handlePreview(i, j)"
                  >
                    <BrowseIcon size="large" />
                  </div>
                </div>
              </div>
              <div class="journal-list-item-meta">
                <span class="journal-list-item-time">{{
                  timeAgo(item.createTime)
                }}</span>
                <span class="journal-list-item-actions">
                  <t-button
                    shape="square"
                    variant="text"
                    @click="handleEdit(item)"
                    ><EditIcon slot="icon"
                  /></t-button>
                  <t-button
                    shape="square"
                    variant="text"
                    @click="handleDelete(item)"
                    ><DeleteIcon slot="icon"
                  /></t-button>
                </span>
              </div>
            </div>
          </template>
        </t-list-item>
      </t-list>
      <NewJournalDialog ref="newJournalDialog"></NewJournalDialog>
    </template>
  </PageView>
</template>

<script>
import { AddIcon, BrowseIcon, EditIcon, DeleteIcon } from "tdesign-icons-vue";
import PageView from "@/layouts/PageView";
import NewJournalDialog from "./components/NewJournalDialog";
import { timeAgo } from "@/utils/datetime.js";

export default {
  name: "JournalList",
  components: {
    AddIcon,
    BrowseIcon,
    EditIcon,
    DeleteIcon,
    PageView,
    NewJournalDialog,
  },
  data() {
    return {
      asyncLoading: "load-more",
      pagination: {
        current: 1,
        pageSize: 10,
        total: 0,
      },
      journalList: [],
      maskVisible: [],
    };
  },
  methods: {
    fetchJournals() {
      this.asyncLoading = "loading";
      this.$request
        .get(
          "journals?pageNum=" +
            this.pagination.current +
            "&pageSize=" +
            this.pagination.pageSize
        )
        .then((res) => {
          this.journalList = res.data.list
            ? this.journalList.concat(...res.data.list)
            : this.journalList;
          if (res.data.current * res.data.pageSize >= res.data.total) {
            this.asyncLoading = "";
          } else {
            this.asyncLoading = "load-more";
          }
          this.pagination = {
            current: res.data.current + 1,
            pageSize: res.data.pageSize,
            total: res.data.total,
          };
          for (let index = 0; index < this.journalList.length; index++) {
            this.$set(
              this.maskVisible,
              index,
              this.journalList[index].images?.map(() => false)
            );
          }
        })
        .catch(() => {
          this.asyncLoading = "load-more";
          this.$message.error("获取日志列表失败");
        });
    },
    onClickNew() {
      this.$refs.newJournalDialog.open();
    },
    timeAgo(time) {
      return timeAgo(time);
    },
    mouseEnter(i, j) {
      this.$set(this.maskVisible[i], j, true);
    },
    mouseLeave(i, j) {
      this.$set(this.maskVisible[i], j, false);
    },
    handlePreview(i, j) {
      this.$viewerApi({
        options: {
          initialViewIndex: j,
        },
        images: this.journalList[i].images,
      });
    },
    handleEdit(journal) {
      this.$message.info("未实现");
      console.log(journal);
    },
    handleDelete(journal) {
      this.$message.info("未实现");
      console.log(journal);
    },
  },
  mounted() {
    this.fetchJournals();
  },
};
</script>
<style lang="less">
.t-list-item__content {
  width: 100%;
}
.journal-list-item-container {
  // display: flex;
  // flex-direction: column;
  width: 100%;
  .journal-list-item-content {
    margin-bottom: 10px;
  }
  .journal-list-item-images {
    display: flex;
    flex-wrap: wrap;
    .journal-list-image-box {
      border-radius: 6px;
      position: relative;
      margin: 0 10px 10px 0;
      width: 100px;
      height: 100px;
      background-position: 50% 50%;
      background-size: cover;
      &:hover {
        transform: scale(1.01);
      }
      .journal-list-image-mask {
        cursor: pointer;
        border-radius: 6px;
        background-color: @text-color-secondary;
        color: @bg-color-container;
        transition: all 0.2s linear;
        will-change: transform;
        opacity: 0;
        display: flex;
        align-items: center;
        justify-content: center;
        position: absolute;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        &:hover {
          opacity: 1;
        }
      }
    }
  }
  .journal-list-item-time {
    font-size: 12px;
    line-height: 24px;
    color: @text-color-placeholder;
  }
  .journal-list-item-actions {
    float: right;
  }
}
</style>