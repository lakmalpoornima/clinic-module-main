<template>
  <CRow>
    <CCol col="12" lg="6">
      <CCard>
        <CCardHeader>
          User id:  {{ $route.params.id }}
         
        </CCardHeader>
        <CCardBody>
          <CDataTable
            striped
            small
            fixed
            :items="visibleData"
            :fields="fields"
          />
        </CCardBody>
        <CCardFooter>
          <CButton color="primary" @click="goBack">Back</CButton>
        </CCardFooter>
      </CCard>
    </CCol>
  </CRow>
</template>

<script>
export default {
  name: 'User',
  inject: ['api'],
  usersData: null,
  beforeRouteEnter(to, from, next) {
    next(vm => {
      vm.usersOpened = from.fullPath.includes('users')
    })
  },

  data () {
    return {
      usersOpened: null
    }
  },

  computed: {
    fields () {
      return [
        { key: 'key', label: this.username, _style: 'width:150px'},
        { key: 'value', label: '', _style: 'width:150px;' }
      ]
    },
    userData () {
      const id = this.$route.params.id
      const user = this.usersData.find((user, index) => index + 1 == id)
      const userDetails = user ? Object.entries(user) : [['id', 'Not found']]
      return userDetails.map(([key, value]) => { return { key, value } })
    },
    visibleData () {
      return this.userData.filter(param => param.key !== 'username')
    },
    username () {
      return this.userData.filter(param => param.key === 'username')[0].value
    }
  },
  created () {
    this.LoadUsers()
  },
  methods: {
    async LoadUsers() {
     const result = await this.api.get("/users")
     this.usersData = result.data
    },
    goBack() {
      this.usersOpened ? this.$router.go(-1) : this.$router.push({path: '/users'})
    }
  },
  // Here have 
  methods:{}
}
</script>
