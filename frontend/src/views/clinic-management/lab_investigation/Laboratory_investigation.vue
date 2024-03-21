<template>
  <div>
   hello
    
  </div>
</template>


<script>
export default {
  data () {
    return {
      warningModal: true,
     
    }
  },


  methods: {
     clickok(){
      this.$router.push('/dashboard')
     }
  }
}
</script>