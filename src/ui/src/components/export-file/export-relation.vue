<template>
  <div class="export-relation">
    <bk-checkbox class="allow-export" v-model="allowExport">{{$t('是否导出关联的模型实例')}}</bk-checkbox>
    <bk-table class="relation-table"
      ref="table"
      :outer-border="false"
      :max-height="$APP.height - 220"
      :data="relations"
      @selection-change="handleSelectionChange">
      <bk-table-column type="selection" :selectable="() => allowExport"></bk-table-column>
      <bk-table-column :label="$t('关联的模型')" prop="model" width="380" :resizable="false">
        <template slot-scope="{ row }">
          <i :class="['model-icon', getModelIcon(row)]"></i>
          {{getModelName(row)}}
        </template>
      </bk-table-column>
      <bk-table-column :label="$t('唯一校验标识')" prop="identification" align="right" :resizable="false">
        <template slot-scope="{ $index }">
          <bk-select class="unique-selector"
            v-model="selected[relationModels[$index]]"
            :clearable="false"
            :disabled="!allowExport">
            <bk-option v-for="uniqueCheck in getRowUniqueChecks($index)"
              :key="uniqueCheck.id"
              :id="uniqueCheck.id"
              :name="getUniqueCheckName(uniqueCheck)">
            </bk-option>
          </bk-select>
        </template>
      </bk-table-column>
      <bk-exception slot="empty" type="empty" scene="part">{{$t('暂无关联模型，无需选择')}}</bk-exception>
    </bk-table>
  </div>
</template>

<script>
  import { ref, computed, watch, reactive, set } from '@vue/composition-api'
  import useState from './state'
  import useRelation from '@/hooks/relation/model'
  import useBatchUniqueCheck from '@/hooks/unique-check/batch'
  import useProperty from '@/hooks/property/property'
  import { mapGetters } from 'vuex'
  export default {
    name: 'export-relation',
    setup() {
      const [state, { setState }] = useState()
      const [relations] = useRelation(state.bk_obj_id)
      const relationModels = computed(() => relations.value.map((item) => {
        const modelId = item.bk_obj_id
        const asstModelId = item.bk_asst_obj_id
        return modelId === state.bk_obj_id.value ? asstModelId : modelId
      }))

      const [properties, { refresh }] = useProperty()
      watch(relationModels, value => refresh({ bk_obj_id: { $in: value } }))
      const [batchUniqueChecks] = useBatchUniqueCheck(relationModels)
      const selected = reactive({})
      watch(batchUniqueChecks, (value) => {
        value.forEach((uniqueChecks) => {
          const preset = uniqueChecks.find(preset => preset.ispre) || uniqueChecks[0]
          set(selected, preset.bk_obj_id, preset.id)
        })
      }, { immediate: true })
      watch(selected, value => setState({ relations: value }), { immediate: true })

      const allowExport = ref(false)
      const handleSelectionChange = () => {}
      return {
        state,
        allowExport,
        handleSelectionChange,
        relations,
        uniqueChecks: batchUniqueChecks,
        properties,
        relationModels,
        selected
      }
    },
    computed: {
      ...mapGetters('objectModelClassify', ['getModelById'])
    },
    watch: {
      allowExport(value) {
        !value && this.$refs.table.clearSelection()
      }
    },
    methods: {
      getRelationModel(relation) {
        const modelId = relation.bk_obj_id === this.state.bk_obj_id.value ? relation.bk_asst_obj_id : relation.bk_obj_id
        return this.getModelById(modelId) || { bk_obj_id: modelId }
      },
      getModelIcon(relation) {
        const model = this.getRelationModel(relation)
        return model.bk_obj_icon || 'icon-cc-default'
      },
      getModelName(relation) {
        const model = this.getRelationModel(relation)
        return model.bk_obj_name || model.bk_obj_id
      },
      getRowUniqueChecks(index) {
        return this.uniqueChecks[index] || []
      },
      getUniqueCheckName({ keys }) {
        const idArray = keys.map(key => key.key_id)
        return idArray.map((id) => {
          const property = this.properties.find(property => property.id === id)
          return property ? property.bk_property_name : `${this.$t('未知属性')}(${id})`
        }).join(' + ')
      }
    }
  }
</script>

<style lang="scss" scoped>
  .export-relation {
    .allow-export {
      display: block;
      margin: 20px 0 0 16px;
    }
  }
  .relation-table {
    margin: 17px 0 0 0;
    .model-icon {
      display: inline-flex;
      justify-content: center;
      align-items: center;
      width: 26px;
      height: 26px;
      background: #f0f1f5;
      border-radius: 50%;
      font-size: 14px;
      margin-right: 10px;
    }
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
      &.is-disabled {
        text-align: right;
        background-color: transparent !important;
        border-color: transparent !important;
        pointer-events: none;
        /deep/ {
          .bk-select-angle {
            display: none;
          }
          .bk-select-name {
            padding: 0;
          }
        }
      }
    }
  }
</style>
