<template v-for="(color, key) in ['success', 'primary', 'info', 'warning', 'danger']">
<div>
 <CRow>
   
    <CCol md="12" > 
          <div class="alertb " v-if="warningModal">
  	            <div class="alertbox">
                 Do you want to delete ?   Department Code : {{messege}}
      
                 <button @click="clickok" color="success" class="abbutton">Accept</button>
                <button type="button" class="close" data-dismiss="alert" aria-label="Close" @click="clickcencel"> 
                  <span aria-hidden="true">&times;</span>
                </button>
     
                </div>
          </div>
     </CCol> 
    
  
  </CRow>
  <CRow>
    

    <CCol col="12">
      <CCard>
        <CCardHeader>
          <div class="d-flex justify-content-between">
               <strong> Department {{this.items.login_id}}</strong> 
                
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
            :items-per-page="10"
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
                <i size="sm"  class="mt-6 " @click="toggleUpdate(item,index)"><CIcon name="cilPencil" class="text-info font-weight-bolder"/></i>
                &nbsp;&nbsp;&nbsp;&nbsp;
                <i size="sm"  @click="toggleDelete(item, index)"><CIcon name="cilTrash" class="text-danger font-weight-bolder"/></i>
              
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
// import New from './Create.vue';
// import 'bootstrap/dist/css/bootstrap.css'
// import 'bootstrap-vue/dist/bootstrap-vue.css'
// import { cilPencil } from '@coreui/icons'
export default {
  // cilPencil
  name: "Users",
  inject: ["api"],
  // props:[items],
  data() {
    return {
      items: null,
      messege:"",
      warningModal: true,
     
      fields: [
        // { key: 'avatar', label: 'Avatar', _classes: 'text-center'}, // No Need for Avatar 
        { key: "department_code", label: "Department code", _classes: "font-weight-bold", },
        { key: "department_name", label: "Department name" },
        { key: "company_code", label: "Company code" },
        { key: "created_on", label: "Create Date"},
        { key: "deleted", label: "Delete"},

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
      for_delete:{ 
        deleted:true,
        deleted_id:"admin"}
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
       this.$router.push('/geo-settings/department')
    // alert('That features not available')
    
    },
    // components: {
    //   New
    // },

   
    async LoadUsers() {
      const result = await this.api.get("/department");
      this.items = result.data;
      console.log(items.department_code)
    },

   

    pageChange(val) {
      this.$router.push({ query: { page: val } });
    },

    toggleDelete (item,index) {
        this.warningModal=true
        this.messege = item.department_code
        this.api.put(`/department_delete/${item.department_code}`,this.for_delete)   
        // this.warningModal=false
      },

    clickok(){
       this.api.put(`/department_delete/${item.department_code}`,this.for_delete)
      try{
          // this.messege = "Deleted successfully";
          
          this.warningModal=false
      }
      catch{
        this.warningModal=false
       }
       
    },

    clickcencel(){
      setTimeout(10000,)
       this.warningModal=false
    },   

    toggleUpdate (item,index) { 
        this.$router.push({ path: `/geo-settings/department/${item.department_code}` });
      },
    load(item){
      loading
    }
    
  },
  
};
</script>
<style>
.alertbox{
  display: flex;
  text-align: center;
 
}
.abbutton{
border: 0.5px solid #1ddcfb;
border-radius: 5px 5px;
margin-left: 15px;
background-color: #2eb85c;
font-weight: bold;
color: white;
}

.alertb{
display: flex;
width: auto;
height: 35px;
text-align: center;
background-color: #c5f7e6;
color: #3781df;
font-weight: bold;
align-items: center;
float: right;
border: 1px solid #1ddcfb;
border-radius: 5px 5px;

 position: absolute;
  z-index: 101;
    top: 5px;
    left: 290px;
    /* right: 0; */
    /* background: #fde073; */
    text-align: center;
    line-height: 2.5;
    overflow: hidden; 
    /* -webkit-box-shadow: 0 0 5px black; */
    /* -moz-box-shadow:    0 0 5px black; */
    /* box-shadow:         0 0 5px black; */
}

.close{
  margin-left: 10px;
  margin-right: 10px;
}
</style>
