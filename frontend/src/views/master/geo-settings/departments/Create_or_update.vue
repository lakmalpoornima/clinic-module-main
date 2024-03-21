<template>

  
<!-- =============================================================Get User Input===================================================== -->
 
    <CRow>

      
      <CCol sm="12">
        <CCard>
          <CCardHeader>
             <div  v-if='this.id'>
            <strong>Update</strong><small>Form</small>
            </div>
            <div v-else>
              <strong>Department</strong> Form
            </div>
            
          </CCardHeader>

          
           <CCardBody>
              <CInput placeholder="Department Code"  valid-feedback="Input is valid." invalid-feedback="Please provide a Department Code."  :is-valid="validator" v-model="posts.department_code">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="Company Code"  valid-feedback="Input is valid." invalid-feedback="Please provide a Company Code." :is-valid="validator" v-model="posts.company_code">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="Department Name" valid-feedback="Input is valid." invalid-feedback="Please provide Department Name."  :is-valid="validator" v-model="posts.department_name">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            
 
          </CCardBody>

           <CCardFooter>
            <div class="d-flex justify-content-center" v-if='this.$route.params.id'>
              <CButton  type="submit" color="info" sm="6" class="font-weight-bold" @click="Update_department()" ><CIcon name="cilCheck"/>&nbsp;Update</CButton> 
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              <CButton  type="Reset" color="info" @click="back" sm="8" class="font-weight-bold" ><CIcon name="cil-cursor"/>&nbsp;Back</CButton>
            </div>

            <div v-else> 
              <CButton  type="submit" color="info" sm="6" class="font-weight-bold" @click="Post_data_for_user()" ><CIcon name="cilCheck"/>&nbsp;Submit</CButton> 
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              <CButton  type="Reset" color="info" @click="back" sm="8" class="font-weight-bold" ><CIcon name="cil-cursor"/>&nbsp;Back</CButton>
            </div>
            
          </CCardFooter>
            
          
        </CCard>
      </CCol>
     
    </CRow>
   
  <!-- </div> -->
   <!-- </CContainer> -->
   
</template>



<script>

import * as Vue from 'vue' 
import axios from 'axios'
import VueAxios from 'vue-axios'; 

import useVuelidate from '@vuelidate/core';
import { Cropper } from 'vue-advanced-cropper';
import 'vue-advanced-cropper/dist/style.css';

import { required,email, integer} from '@vuelidate/validators';
// import fileupload from 'public/img/avatars';
import { reactive,computed} from 'vue';

export default{
  name:"Create",//Component Name
   inject: ["api"],
 
   

  data(){ //data part
  // console.log("Goo",imageData)
    return{
      visibleLiveDemo: false,
      departmets:[],
     
      posts:{
          department_code: null,
          company_code:null,
          department_name:null
         
      },
      
      coordinates: {
            width: 0,
            height: 0,
            left: 0,
            top: 0,
          },
    }
   
  },

  props:{
    id:{ type:integer, required:true}
  },
   validations () {
    return {
      company_code: { required }, // Matches this.firstName
      department_name: { required }, // Matches this.lastName
      department_code: { required} // Matches this.contact.email
      }
    },
  // this for edit 
   created: function(){
            this.getItem();
    },

    // Getdepartment(){
    //    this.find_department();
    // },


// Methods 
  methods:{
    find_department(){
      let department_code = this.$route.params.id
     
       this.api.get(`http://127.0.0.1:9085/api/department/${department_code}`)
        .then((response)=>{
            this.posts=response.data;
           console.log(this.user1," For find a user")
        });

        
      
      
    },
    // call the funtion 
    refreshData(){
        this.api.get("/username")
        .then((response)=>{
            this.user=response.data;
          //  console.log(this.user," For find a user")
        });
        return this.user
    },
// validator
     validator (val) {
      return val ? val.length >= 2 : false
    },

//  Send the create record 
    Post_data_for_user(){
      
      this.api.post("/department",this.posts)
      .then((result)=>{
        this.result = result.data
        alert("Saved");
        console.log("result.data",this.result)
      })

     
      .catch(error =>console.log(error))
      this.v$.$validations
      if (!this.v$.$error){
          alert('Form successfully submitted')
      }else{
          alert('Form failed validation')
      }
    },
// Update
      Update_department(){
        let department_code = this.$route.params.id
        this.api.put(`/department/${department_code}`,this.posts)
      .then((result)=>{
        this.result = result.data
        alert("Saved");
        console.log("result.data",this.result)
      })

     
      .catch(error =>console.log(error))
      },
// Delete
    Delete(){
       
    },
    back:function(){
     
       this.$router.push('/geo-settings/departments')

    },


// Select and upload Profile Image



    
   
    
  // for show entered Password string  
   

      
},
  
  mounted:function(){
    this.refreshData();
    this.find_department();
  },
 
}
</script>

<style scoped>
.profile{
  width:200px;
  height:200px;
  align-self: center;
  padding: 20px;

}
/* #file {
 display: none;
} */
.image{
  position: absolute;
  text-align: center;
  bottom: 25px;
  transform: translateX(-50% );
  height: 40px;
  width: 10%;
  left: 50%;
  line-height: 50px;
}



.fileup {
 display: none;
}

.profiler{
 height: 200px;
    width: 200px;
    /* position:absolute; */
    top: 30%;
    left: -33%;
    transform: translate(40%,-10%);
    border-radius: 50%;
    overflow: hidden;
    /* border: 1px solid grey; */
  
}
.preview-image{
  width: 200px;
  height: 200px;

}

.base-image-input {
  display: block;
  width: 200px;
  height: 200px;
  cursor: pointer;
  background-size: cover;
  background-position: center center;
}.placeholder {
  background: #F0F0F0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #333;
  font-size: 18px;
  font-family: Helvetica;
}.placeholder:hover {
  background: #E0E0E0;
}.file-input {
  display: none;
}
</style>
