<template>
  <div class="import-state">
    <template v-if="state === 'pending'">
      <img class="state-loading" src="../../assets/images/icon/loading.svg" alt="loading">
      <p class="text">{{$t('数据导入中')}}</p>
    </template>
    <template v-else-if="state === 'success'">
      <i class="state-success bk-icon icon-check-circle-shape"></i>
      <p class="text">{{$t('导入成功')}}</p>
      <bk-button class="mt20" theme="default" @click="closeImport">{{$t('关闭')}}</bk-button>
    </template>
    <template v-else-if="state === 'error'">
      <i class="state-error bk-icon icon-close-circle-shape"></i>
      <p class="text">{{$t('导入失败')}}</p>
      <bk-button class="mt20" theme="default" @click="reset">{{$t('重新导入')}}</bk-button>
    </template>
  </div>
</template>

<script>
  import useFile from './file'
  import useImport from './index'
  import useStep from './step'
  export default {
    name: 'import-state',
    setup() {
      const [{ state }, { clear: clearFile }] = useFile()
      const [, { close: closeImport }] = useImport()
      const [, { reset: resetStep }] = useStep()
      const reset = () => {
        resetStep()
        clearFile()
      }
      return {
        state,
        reset,
        closeImport
      }
    }
  }
</script>

<style lang="scss" scoped>
  .import-state {
    display: flex;
    flex-direction: column;
    align-items: center;
    .state-loading {
      margin: 20px 0 16px 0;
      width: 60px;
    }
    .text {
      font-size: 22px;
      color: #313238;
      line-height: 30px;
    }
    .state-success,
    .state-error {
      font-size: 60px;
      margin: 20px 0 16px 0;
    }
    .state-success {
      color: $successColor;
    }
    .state-error {
      color: $dangerColor;
    }
  }
</style>
