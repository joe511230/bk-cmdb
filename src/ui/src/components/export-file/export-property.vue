<template>
  <div class="property-selector">
    <div class="filter">
      <bk-input v-model.trim="keyword" :placeholder="$t('请输入字段名称')"></bk-input>
    </div>
    <div class="group-list"
      v-for="{ group, properties } in groupedPropertyies"
      v-show="properties.length"
      :key="group.id">
      <div class="group-header">
        <label class="group-label">{{group.bk_group_name}}</label>
        <bk-checkbox class="group-checkbox"
          :checked="isAllSelected(properties)"
          @change="setAllSelection(properties, ...arguments)">
          {{$t('全选')}}
        </bk-checkbox>
      </div>
      <ul class="property-list">
        <li class="property-item"
          v-for="property in properties"
          :key="property.id">
          <bk-checkbox class="property-checkbox"
            :title="property.bk_property_name"
            :checked="isSelected(property)"
            @change="setSelection(property, ...arguments)">
            {{property.bk_property_name}}
          </bk-checkbox>
        </li>
      </ul>
    </div>
    <bk-exception type="search-empty" scene="part" v-show="!matchedProperties.length"></bk-exception>
  </div>
</template>

<script>
  import { ref, watch } from '@vue/composition-api'
  import useFilter from '@/hooks/utils/filter'
  import useGroupProperty from '@/hooks/utils/group-property'
  import useProperty from '@/hooks/property/property'
  import useGroup from '@/hooks/property/group'
  import useState from './state'
  export default {
    name: 'import-property',
    setup() {
      const [exportState, { setState }] = useState()
      const [properties] = useProperty({
        bk_obj_id: exportState.bk_obj_id.value,
        bk_biz_id: exportState.bk_biz_id.value
      })
      const [groups] = useGroup({
        bk_obj_id: exportState.bk_obj_id.value,
        bk_biz_id: exportState.bk_biz_id.value
      })
      const keyword = ref('')
      const [matchedProperties] = useFilter({
        list: properties,
        keyword,
        target: 'bk_property_name',
        available: exportState.available.value
      })
      const groupedPropertyies = useGroupProperty(groups, matchedProperties)

      const selection = ref([])
      const setSelection = (item, selected) => {
        if (selected) {
          selection.value.push(item)
          return
        }
        const index = selection.value.indexOf(item)
        index > -1 && selection.value.splice(index, 1)
      }
      const setAllSelection = (properties, selected) => {
        if (selected) {
          selection.value = [...new Set([...selection.value, ...properties])]
        } else {
          selection.value = selection.value.filter(property => !properties.includes(property))
        }
      }
      const isSelected = property => selection.value.includes(property)
      const isAllSelected = properties => properties.every(property => selection.value.includes(property))
      watch(selection, value => setState({ selection: value }))

      return {
        keyword,
        matchedProperties,
        groupedPropertyies,
        selection,
        setSelection,
        isSelected,
        setAllSelection,
        isAllSelected
      }
    }
  }
</script>

<style lang="scss" scoped>
  .property-selector {
    padding: 20px 0 0 0;
    .filter {
      width: 340px;
    }
  }
  .group-list {
    margin: 25px 0 0 0;
    .group-header {
      display: flex;
      align-items: center;
      padding: 0 0 10px;
      .group-label {
        flex: 1;
        font-weight: 700;
        line-height: 20px;
        &:before {
          content: "";
          display: inline-block;
          vertical-align: top;
          width: 4px;
          height: 16px;
          background: #dcdee5;
          margin: 2px 9px 0 0;
        }
      }
      .group-checkbox {
        margin-left: auto;
      }
    }
  }
  .property-list {
    display: flex;
    flex-wrap: wrap;
    .property-item {
      width: 33%;
      line-height: 34px;
      .property-checkbox {
        display: inline-flex;
        vertical-align: middle;
        /deep/ {
          .bk-checkbox {
            width: 16px;
          }
          .bk-checkbox-text {
            max-width: 160px;
            @include ellipsis;
          }
        }
      }
    }
  }
</style>
