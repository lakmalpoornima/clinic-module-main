<template v-for="(color, key) in ['success', 'primary', 'info', 'warning', 'danger']">
<div>
 <CRow>
    <!-- <CBreadcrumb :items="items"/> -->
    <!-- <CCol col>
      <CCard>
        <CCardHeader>
         
        </CCardHeader>
        
      </CCard>
    </CCol> -->
  </CRow>
  <CRow>
    

    <CCol col="12">
      <CCard>
        <CCardHeader>
          <div class="d-flex justify-content-between">
               <strong>USERS{{this.items.login_id}}</strong> 
                <!-- <CButton class="font-weight-bold p-20" @click="create_new"><CIcon name="cilUser" color ="success"/>Create New User</CButton> -->
                <i class="font-weight-bold p-20" @click="create_new"><CIcon name="cil-user-follow" class="text-success font-weight-bolder"/></i>
             </div>
          

           <!-- <div class="d-flex justify-content-around"><CButton color="success" class="font-weight-bold" @click="create_new">Create New User</CButton></div> -->
        </CCardHeader>
        <CCardBody>
          <CButtonGroup class="button">
          
          </CButtonGroup>
           
          <CDataTable
            header
            hover
            striped
            :items="items"
            :fields="fields"
            :items-per-page="25"
            clickable-rows
            :active-page="activePage"
            @row-clicked="rowClicked"
            :pagination="{ doubleArrows: false, align: 'center' }"
            @page-change="pageChange"
            column-filter
            table-filter
            sorter
            items-per-page-select
            @loadstart="load"
            >
         
           <template #show_details="{item, index}">
            <td class="py-10">
              <!-- <CButtonGroup >
                <CButton size="sm"  class="mt-6 " @click="toggleUpdate(item, index)"><CIcon name="cilPencil"/>Edit</CButton>
                <CButton size="sm"  @click="toggleDelete(item, index)"><CIcon name="cilTrash"/>Delete</CButton>
              </CButtonGroup> -->
               <CButtonGroup >
                <i size="sm"  class="mt-6 " @click="toggleUpdate(item, index)"><CIcon name="cilPencil" class="text-info font-weight-bolder"/></i>
                &nbsp;&nbsp;&nbsp;&nbsp;
                <i size="sm"  @click="toggleDelete(item, index)"><CIcon name="cilTrash" class="text-danger font-weight-bolder"/></i>
                <!-- <i size="sm"  @click="warningModal = true"><CIcon name="cilTrash" class="text-danger font-weight-bolder"/>
                   <CModal title="Modal title" color="warning" :show.sync="warningModal" > </CModal>
                </i>  -->
              </CButtonGroup>
              
            </td>
          </template>
          </CDataTable>
        </CCardBody>
      </CCard>
    </CCol>
  </CRow>
  <!-- </CContainer> -->
</div>
</template>

<script>
import New from './Create.vue';

export default {
  // cilPencil
  name: "Users",
  inject: ["api"],
  // props:[items],
  data() {
    return {
      items: null,
      fields: [
        // { key: 'avatar', label: 'Avatar', _classes: 'text-center'}, // No Need for Avatar 
        { key: "login_id", label: "ID", _classes: "font-weight-bold", },
        { key: "user_name", label: "Name"},
        { key: "create_date", label: "Create Date"},
        { key: "email", label: "Email" },
        {
            key: 'show_details',
            label: '',
            _style: { width: '1%' },
            filter: false,
            sorter: false,
            _props: { color: 'primary', className: 'fw-semibold'}
          },
      ],
      activePage: 1,
      collapseDuration: 4,
      details: [],
      warningModal: false,
      size:'sm',
    };
  },

  watch: {
    $route: {
      immediate: true,
      handler(route) {
        if (route.query && route.query.page) {
          this.activePage = Number(route.query.page);
        }
      }
    }
  },
  created() {
    this.LoadUsers();
    // this.LoadAvatar();
  },
  methods: {
    
    create_new:function(){
       this.$router.push('/system-security/UserAction')
    
    },
    components: {
      New
    },

   
    async LoadUsers() {
      const result = await this.api.get("/users");
      this.items = result.data;
    },

    // rowClicked(item, index) {
    //   this.$router.push({ path: `users/${index + 1}` });
    // },

    pageChange(val) {
      this.$router.push({ query: { page: val } });
    },

    toggleDelete (item) {
        //  this.items.pop()
        if(!confirm("Are you sure")){
          return
        }
        // axios.delete("/users"+this.id) // This delete function has work fine
        
      },

    toggleUpdate (item,index) {
      // this.$router.push({ path: `user/${index + 1}` });
      console.log("get print id", )
      this.$router.push({ path: `user/${this.items.id}` });
      },
    load(item){
      loading
    }
    
  },

  

  
  
};
</script>
