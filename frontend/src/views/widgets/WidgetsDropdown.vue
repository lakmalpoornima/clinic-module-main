<template>
  <CRow>
    <CCol sm="6" lg="3">
      <CWidgetDropdown v-for="tile in tiles" :key="tile.header" color="gradient-primary" :header="tile.header" :text="tile.data">
      </CWidgetDropdown>
    </CCol>
  </CRow>
</template>

<script>
import { CChartLineSimple, CChartBarSimple } from '../charts/index.js'

export default {
  name: 'WidgetsDropdown',
  components: { CChartLineSimple, CChartBarSimple },
  props: {
    tiles: Array
  }
}
</script>
