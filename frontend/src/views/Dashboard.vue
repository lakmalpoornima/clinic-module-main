<template>
  <div>
    <CRow>
      <CCol class="" :key="tile.header" v-for="tile in tiles">
      <CWidgetIcon
        :header="tile.data"
        :text="tile.header"
        color="info"
      >
        <CIcon name="cil-people" height="36"/>
      </CWidgetIcon>
      </CCol>
    </CRow>
  </div>
</template>

<script>


export default {
  name: 'Dashboard',
  inject: ["api"],
  data () {
    return {
      tiles: []
    }
  },
  methods: {
    async GetDashboardCount(){
      const result = await this.api.get('/dashboard/GetDashboardCount')
      let header = result.data
      let item = {}
      item.header = "Patient Count Waiting In Clinic Receiption"
      item.data = header.ReceptionCount.toString()
      this.tiles.push(item)
      item = {}
      item.header = "Current Patient Count In Clinic"
      item.data = header.InCount.toString()
      this.tiles.push(item)
      item = {}
      item.header = "Total Patient Servered Count"
      item.data = header.ServedCount.toString()
      this.tiles.push(item)
      item = {}
      item.header = "Patient Count Servered From Me"
      item.data = header.ServedFromMe.toString()
      this.tiles.push(item)
    }
  },
  created() {
    this.GetDashboardCount()
  }
}
</script>
