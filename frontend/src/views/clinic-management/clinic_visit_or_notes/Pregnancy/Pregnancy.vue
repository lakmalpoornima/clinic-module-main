<template>
  <div>
   <CModal
      color="warning"
	    :no-close-on-backdrop="true"
      :show.sync="warningModal"
      title="Apologize....!"
    >
		Without Select a <strong>  Patient </strong> can not be retrieve Pregnancy-diabetes.
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


  created() {
    this.delay();
  },


  methods: {
     delay(){
       setTimeout(function () {
            window.location.href = "#/doctor-queue-or-portal";}, 1000);
     }
  }
}
</script>