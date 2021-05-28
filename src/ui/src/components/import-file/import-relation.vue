<template>
  <div class="import-relation">
    <bk-checkbox class="allow-import" v-model="allowImport">{{$t('是否导入关联的模型实例')}}</bk-checkbox>
    <bk-table class="relation-table"
      ref="table"
      :outer-border="false"
      :max-height="$APP.height - 220"
      :data="list"
      @selection-change="handleSelectionChange">
      <bk-table-column type="selection" :selectable="() => allowImport"></bk-table-column>
      <bk-table-column :label="$t('关联的模型')" prop="model" width="380" :resizable="false"></bk-table-column>
      <bk-table-column :label="$t('唯一校验标识')" prop="identification" align="right" :resizable="false">
        <template slot-scope="{ row }">
          <bk-select class="unique-selector" v-model="row.model" :clearable="false">
            <bk-option id="model" name="model"></bk-option>
          </bk-select>
        </template>
      </bk-table-column>
      <bk-exception slot="empty" type="empty" scene="part">{{$t('暂无关联模型，无需选择')}}</bk-exception>
    </bk-table>
    <div class="options">
      <bk-button theme="default" @click="previousStep">{{$t('上一步')}}</bk-button>
      <bk-button theme="primary" class="ml10" @click="startImport">{{$t('开始导入')}}</bk-button>
      <bk-button theme="default" class="ml10" @click="closeImport">{{$t('取消')}}</bk-button>
    </div>
  </div>
</template>

<script>
  import useStep from './step'
  import useImport from './index'
  import useFile from './file'
  export default {
    name: 'import-relation',
    setup() {
      const [currentStep, { previous: previousStep }] = useStep()
      const [importState, { close: closeImport }] = useImport()
      const [{ file }, { setState: setFileState, setError: setFileError }] = useFile()
      return {
        currentStep,
        previousStep,
        closeImport,
        importState,
        file,
        setFileState,
        setFileError
      }
    },
    data() {
      return {
        list: [],
        selection: [],
        allowImport: false
      }
    },
    watch: {
      allowImport(allowImport) {
        if (!allowImport) {
          this.$refs.table.clearSelection()
        }
      }
    },
    methods: {
      handleSelectionChange(selection) {
        this.selection = selection
      },
      async startImport() {
        try {
          this.setFileState('pending')
          await this.importState.submit({
            file: this.file,
            step: this.currentStep
          })
          this.setFileState('success')
        } catch (error) {
          console.error(error)
          this.setFileState('error')
          this.setFileError(error)
        }
      }
    }
  }
</script>

<style lang="scss" scoped>
  .import-relation {
    .allow-import {
      display: block;
      margin: 20px 0 0 16px;
    }
  }
  .relation-table {
    margin: 17px 0 0 0;
    /deep/ .bk-table-row {
      &.hover-row > td {
        background-color: #fff;
      }
      &:not(.hover-row){
        .unique-selector:not(.is-focus) {
          background-color: transparent;
          text-align: right;
          .bk-select-angle {
            display: none;
          }
          .bk-select-name {
            padding: 0;
          }
        }
      }
    }
    .unique-selector {
      width: 100%;
      background: #f0f1f5;
      border-color: transparent;
      text-align: left;
      &.is-focus {
        box-shadow: none;
      }
    }
  }
  .options {
    margin: 20px 0 0 0;
    display: flex;
  }
</style>
