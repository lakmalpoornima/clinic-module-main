<template>
  <CFooter :fixed="false">
    <div>
      <a href="https://21ccare.com" target="_blank">Systolic</a>
      <span class="ml-1">&copy; {{new Date().getFullYear()}} Version: {{ this.$Version }} .</span>
    </div>
    <div class="mfs-auto">
      <span class="mr-1" target="_blank">Powered by</span>
      <a href="https://21ccare.com">Systolic</a>
    </div>
  </CFooter>
</template>

<script>
export default {
  name: 'TheFooter'
}
</script>
