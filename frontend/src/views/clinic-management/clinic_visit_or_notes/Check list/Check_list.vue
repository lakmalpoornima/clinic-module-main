<template>
  <div>
   <CCol md="12" > 
          <div class="alertb " v-if="warningModal">
  	            <div class="alertbox">
                 Without select a patient can not reach out the page 
      
                 <!-- <button @click="delay" color="success" class="abbutton">Accept</button> -->
                <button type="button" class="close" data-dismiss="alert" aria-label="Close" @click="clickcencel"> 
                  <span aria-hidden="true">&times;</span>
                </button>
     
                </div>
          </div>
     </CCol> 
    
  </div>
</template>


<script>
export default {
  data () {
    return {
      warningModal: true,
     
    }
  },
  
  created() {
    this.delay();
  },

  methods: {
     
     delay(){
       setTimeout(function () {
            window.location.href = "#/doctor-queue-or-portal"; 
            // this.$router.push('/Clinic Management/doctor-queue-or-portal')
                              }, 1000);
     }

  }
}
</script>