<template>
  <div>
   <CModal
      color="warning"
	    :no-close-on-backdrop="true"
      :show.sync="warningModal"
      title="Apologize....!"
    >
		
        Without Select a <strong>  Patient </strong> can not be retrieve Check List.
	<p>Thank You</p>
      <template #footer>
        <!-- <CButton @click="clickok" color="danger">Discard</CButton> -->
        <CButton @click="clickok" color="success">Accept</CButton>
      </template>
    
    </CModal>
    
  </div>
</template>


<script>
export default {
  data () {
    return {
      warningModal: true,
     
    }
  },


   mounted() {
    this.delay();
    
  },


  methods: {
     delay(){
       setTimeout(function () {
            window.location.href = "#/doctor-queue-or-portal"; 
            // this.$router.push('/Clinic Management/doctor-queue-or-portal')
                              }, 2000);
     }
  }
}
</script>