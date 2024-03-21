<template v-for="(color, key) in ['success', 'primary', 'info', 'warning', 'danger']">
<div>
 <CRow>
  <CCol md="12" > 
          <div class="alertb " v-if="warningModal">
  	            <div class="alertbox">
                 Do you want to delete ?  Company Code : {{messege}}
      
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
               <strong> Company info {{this.items.login_id}}</strong> 
                <i class="font-weight-bold p-20" @click="create_new"><CIcon name="cil-user-follow" class="text-success font-weight-bolder"/></i>
             </div>
          

           
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
             
               <CButtonGroup >
                <i size="sm"  class="mt-6 " @click="toggleUpdate(item, index)"><CIcon name="cilPencil" class="text-info font-weight-bolder"/></i>
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
        { key: "company_name", label: "Company Name", _classes: "font-weight-bold", },
        { key: "company_code", label: "Company code"},
        
        { key: "email", label: "Email" },
                
        { key: "address", label: "Address" },
        { key: "telephone", label: "Telephone" },
        { key: "fax", label: "Fax" },
        { key: "url", label: "URL" },
        { key: "create_date", label: "Create Date"},
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
       this.$router.push('/geo-settings/company')
    
    },
   

   
    async LoadUsers() {
      const result = await this.api.get("/company_info");
      this.items = result.data;
    },

    
    pageChange(val) {
      this.$router.push({ query: { page: val } });
    },

    toggleDelete (item,index) {
        this.warningModal=true
        this.messege = item.company_code
         this.api.put(`/company_info_delete/${item.company_code}`,this.for_delete)
      //   .then((result)=>{
      //   this.result = result.data
      //   console.log("Data",this.result,"hoh",Company_info)
      //   alert(this.result,"hoh",Company_info);
      // })
     
        
      },

    toggleUpdate (item,index) {
     
      console.log("get print id")
      this.$router.push({ path: `/geo-settings/company/${item.company_code}` });
      },


    load(item){
      loading
    },

    // here to the get rejected to modal
    clickcencel(){
       setTimeout(10000,)
       this.warningModal=false
    },

    // here to the 
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

    
  },
  
};
</script>
