<template>

<div>
  
 

    <CRow sm="12">
      <CCol>
        <CCard >
          <CCardHeader>
            <strong>Prescription Note</strong> Form
          </CCardHeader>

          <CCardBody  >
            <CForm>
            
               <CInput
                label="ID:"
                valid-feedback="Input is valid."
                invalid-feedback="Please provide at least 4 characters."
                value="Valid value"
                :is-valid="validator"
              />

              <CRow md="12">
                  <CCol  md="4"> 
                  <CInput
                label="P:"
                valid-feedback="Input is valid."
                invalid-feedback="Please provide at least 4 characters."
                value="Valid value"
                :is-valid="validator"
              />
                  </CCol>
                  <CCol  md="4">
                  <CInput
                label="C:"
                valid-feedback="Input is valid."
                invalid-feedback="Please provide at least 4 characters."
                value="Valid value"
                :is-valid="validator"
              />
                  </CCol>
                  <CCol  md="4">
                 <CInput
                label="Obstretic"
                valid-feedback="Input is valid."
                invalid-feedback="Please provide at least 4 characters."
                value="Valid value"
                :is-valid="validator"
              />
                  </CCol>
              </CRow>
 
<CTextarea
                label="Notes"
                valid-feedback="Input is valid."
                invalid-feedback="Please provide at least 4 characters."
                value="Valid value"
                :is-valid="validator"
              />
               
               

            </CForm>
                      

          </CCardBody>

        </CCard>
      </CCol>
      
    </CRow>

  </div>


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
  // setup() {
  //   const state = reactive({
  //     company_code:'',
  //     company_name:'',
  //     email:'',
  //   })
  // },
   

  data(){ //data part
  // console.log("Goo",imageData)
    return{
      //  imageData: null,
      //  progress :0,
      //  getlocal: null,
     
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
      
      // coordinates: {
      //       width: 0,
      //       height: 0,
      //       left: 0,
      //       top: 0,
      //     },
    }
   
  },
  props:{
    id:{ type:integer, required:true}
  },
   validations () {
    return {
      company_code: { required }, // Matches this.firstName
      company_name: { required }, // Matches this.lastName
      email: { required,email} // Matches this.contact.email
      }
    },
  // this for edit 
   created: function(){
            this.getItem();
    },


// Methods 
  methods:{
    // call the funtion 
    // refreshData(){
    //     this.api.get("users/")
    //     .then((response)=>{
    //         this.user=response.data;
    //         // this.departmentsWithoutFilter=response.data;
    //     });
    // },

     validator (val) {
      return val ? val.length >= 4 : false
    },

//  Send the create record 
    Post_data_for_user(){
      
      this.api.post("/company_info",this.posts)
      .then((result)=>{
        this.result = result.data
        alert(result.data);
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


    back:function(){
      //  console.log(getlocal,'IDD')
       this.$router.push('/dashboard')
    
    },

   

// Select and upload Profile Image

// onSelectFile () {
//   const input = this.$refs.fileInput
//   const files = input.files

//   console.log(files,"yoooo")
  
//   if (files && files[0]) {
//     const reader = new FileReader

//     reader.onload = e => {
//     this.imageData = e.target.result
//     // have to store in localstorage 
//     let getlocal = this.imageData
//     localStorage.setItem("image",getlocal)
//      return getlocal
//     }  

//     reader.readAsDataURL(files[0])
//     this.$emit('input', files[0])
//   }
//   else
//   {
//     alert('Sorry, FileReader API not supported')
//   }
  

  
//   },

    
    // updatedata(up){
    //   this.api.put("/user/"+this.id,{
    //     userid:this.userid,
    //     email:this.email,
    //     name:this.name,
    //   })
    //   .then((response)=>{alert(response.data)});
    
    // },
    
  // for show entered Password string  
    // show_password:function(){
    //       // const togglePassword = document.querySelector('#togglePassword');
    //       const password = document.querySelector('#password');
    //       const type = password.getAttribute('type') === 'password' ? 'text' : 'password';
    //       password.setAttribute('type', type);
    //         // toggle the eye / eye slash icon
    //       this.classList.toggle('bi-eye');
    //   },

      // change({ coordinates, canvas }) {
			// console.log(coordinates, canvas);
		  // },
},
  
  mounted:function(){
    this.refreshData();
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
.radiobox{
  display: flex;
  /* padding-left: 2%; */
  margin-left: 10px;
}
.margin_set{
  margin: 10px;
}
.form-check-input{
   padding-left: -100%; 
   background-color: aqua;
}
</style>
