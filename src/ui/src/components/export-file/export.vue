<template>
  <cmdb-sticky-layout class="export">
    <bk-steps class="export-steps" :steps="steps" :cur-step="currentStep"></bk-steps>
    <keep-alive>
      <component :is="stepComponent"></component>
    </keep-alive>
    <div :class="['options', { 'is-sticky': sticky }]" slot="footer" slot-scope="{ sticky }">
      <template v-if="currentStep === 1">
        <bk-button class="mr10" theme="primary"
          :disabled="!selection.length"
          @click="nextStep">
          {{$t('下一步')}}
        </bk-button>
      </template>
      <template v-else>
        <bk-button class="mr10" theme="default" @click="previousStep">{{$t('上一步')}}</bk-button>
        <bk-button class="mr10" theme="primary" @click="startExport">{{$t('开始导出')}}</bk-button>
      </template>
      <bk-button theme="default">{{$t('取消')}}</bk-button>
    </div>
  </cmdb-sticky-layout>
</template>

<script>
  import exportProperty from './export-property'
  import exportRelation from './export-relation'
  import useState from './state'
  import { computed } from '@vue/composition-api'
  export default {
    name: 'export-file',
    components: {
      [exportProperty.name]: exportProperty,
      [exportRelation.name]: exportRelation
    },
    setup() {
      const [{ step: currentStep, selection, relations, submit }, { setState }] = useState()
      const nextStep = () => setState({ step: currentStep.value + 1 })
      const previousStep = () => setState({ step: currentStep.value - 1 })
      const stepComponent = computed(() => ({ 1: exportProperty.name, 2: exportRelation.name }[currentStep.value]))
      return { selection, relations, currentStep, nextStep, previousStep, stepComponent, setState, submit }
    },
    data() {
      return {
        steps: [{ title: this.$t('选择字段'), icon: 1 }, { title: this.$t('选择关联模型'), icon: 2 }]
      }
    },
    methods: {
      async startExport() {
        try {
          this.setState({ status: 'pending' })
          await this.submit({ selection: this.selection, relations: this.relations })
          this.setState({ status: 'success' })
        } catch (error) {
          this.setState({ status: error })
          console.error(error)
        }
      }
    }
  }
</script>

<style lang="scss" scoped>
  .export {
    height: 100%;
    padding: 20px 28px 0;
    @include scrollbar-y;
    .export-steps {
      width: 350px;
      margin: 0 auto;
    }
    .options {
      height: 50px;
      width: calc(100% + 56px);
      margin: 20px 0 0 -28px;
      padding: 0 28px;
      display: flex;
      align-items: center;
      background-color: #fff;
      &.is-sticky {
        border-top: 1px solid $borderColor;
      }
    }
  }
</style>
