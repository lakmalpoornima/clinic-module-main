<template>

  
<!-- =============================================================Get User Input===================================================== -->
 
    <CRow>

      
      <CCol sm="12" md="12">
        <CCard>
          <CCardHeader>
             <div  v-if='this.$route.params.id'>
            <strong>Update</strong><small>Form</small>
            </div>
            <div v-else>
              <strong>Company Profile</strong> Form
            </div>
            
          </CCardHeader>

          
           <CCardBody>
            <CInput placeholder="Company Code" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide a Company Code."
                
                :is-valid="Company_codeIsValid" v-model="posts.company_code" >
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="Company Name" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide Company Name."
                
                :is-valid="Company_nameIsValid" v-model="posts.company_name">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="Address" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide a an Address."
                
                :is-valid="AddressIsValid" v-model="posts.address">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="Telephone" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide Telephone."
                
                :is-valid="TelephoneIsValid" v-model="posts.telephone">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="E-mail" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide a E-mail."
               
                :is-valid="EmailIsValid" v-model="posts.email">
              <template #prepend-content><CIcon name="cil-user"/></template>
            </CInput>

            <CInput placeholder="Fax" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide a Fax."
                v-model="posts.fax"
                :is-valid="FaxIsValid">
              <template #prepend-content><CIcon name="cil-user" /></template>
            </CInput>

            <CInput placeholder="URL" 
                valid-feedback="Input is valid."
                invalid-feedback="Please provide a Url."
                :is-valid="UrlIsvalid"
                v-model="posts.url"
                >
              <template #prepend-content><CIcon name="cil-user" /></template>
            </CInput>

            
          </CCardBody>

           <CCardFooter>
            <div class="d-flex justify-content-center" v-if='this.$route.params.id'>
              <CButton  type="submit" color="info" sm="6" class="font-weight-bold" @click="Company_info_update"><CIcon name="cilCheck"/>&nbsp;Update</CButton> 
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              <CButton  type="Reset" color="info" @click="back" sm="8" class="font-weight-bold" ><CIcon name="cil-cursor"/>&nbsp;Back</CButton>
            </div>

            <div v-else> 
              <CButton  type="submit" color="info" sm="6" class="font-weight-bold" @click="Post_data_for_company" ><CIcon name="cilCheck"/>&nbsp;Submit</CButton> 
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              <CButton  type="Reset" color="info" @click="back" sm="8" class="font-weight-bold" ><CIcon name="cil-cursor"/>&nbsp;Back</CButton>
            </div>
            
          </CCardFooter>
            
          
        </CCard>
      </CCol>
    </CRow>

   
  
</template>



<script>

import * as Vue from 'vue' 
import axios from 'axios'
import VueAxios from 'vue-axios'; 

import useVuelidate from '@vuelidate/core';
import { Cropper } from 'vue-advanced-cropper';
import 'vue-advanced-cropper/dist/style.css';

import { required,email, integer} from '@vuelidate/validators';

import { reactive,computed} from 'vue';

export default{
  name:"Create",
   inject: ["api"],
  

  data(){
    return{
     
     
      posts:{
          company_code:"",
          company_name:"",
          address:"",
          telephone:"",
          email:"",
          create_id:"",
          fax:"",
          url:""
      },
      
    }
   
  },
  props:{
    id:{ type:integer, required:true}
  },
  

    // connection (){
    //   this.LoadCompanyinfo() 
    // },

  // this for edit 
  //  created: function(){
  //           this.getItem();
  //   },


// Methods 
computed:{
 
    Company_codeIsValid(){
      return !!this.posts.company_code
    },
    Company_nameIsValid(){
      return !!this.posts.company_name
    },
    AddressIsValid(){
      return !!this.posts.address
    },
    TelephoneIsValid(){
      return !!this.posts.telephone
    },
    EmailIsValid(){
      return !!this.posts.email
    },

    FaxIsValid(){
      return !!this.posts.fax
    },
    UrlIsvalid(){
      return !! this.posts.url
    },
    FormsubmitCI(){
      return this.Company_codeIsValid && this.Company_nameIsValid &&  this.TelephoneIsValid && this.EmailIsValid && this.FaxIsValid && this.UrlIsvalid && this.AddressIsValid
    }
},
  methods:{
    

 async LoadCompanyinfo() {
      let Company_info = this.$route.params.id
      const result = await this.api.get(`/company_info/${Company_info}`);
      this.posts = result.data;
      console.log("gooo",this.posts)
    },
//  Send the create record 
    Post_data_for_company(){
      
      this.api.post("/company_info",this.posts)
      .then((result)=>{
        this.result = result.data
        alert(this.result);
      })

     
      .catch(error =>console.log(error))
      this.v$.$validations
      if (!this.v$.$error){
          alert('Form successfully submitted')
      }else{
          alert('Form failed validation')
      }

      e.preventDefault();
      alert(result.data);
      console.log("home")
    },

    Company_info_update(){
      let Company_info = this.$route.params.id
      if(this.FormsubmitCI){
        this.api.put(`/company_info/${Company_info}`,this.posts)
        .then((result)=>{
        this.result = result.data
        console.log("Data",this.result,"hoh",Company_info)
        alert(this.result,"hoh",Company_info);
      })

     
      .catch(error =>console.log(error))
        alert("Update function");
      }
        else{
          alert("Unable to Update");
        }
    },


    back:function(){
       this.$router.push('/geo-settings/company info')
    },

},
  
  mounted:function(){
  
    this.LoadCompanyinfo();
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
