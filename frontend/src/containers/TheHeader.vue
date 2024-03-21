<template>
  <CHeader with-subheader>
    <CToggler
      in-header
      class="ml-3 d-lg-none"
      @click="$store.commit('toggleSidebarMobile')"
    />
    <CToggler
      in-header
      class="ml-3 d-md-down-none"
      @click="$store.commit('toggleSidebarDesktop')"
    />
    <CHeaderBrand class="mx-auto d-lg-none" to="/">
    </CHeaderBrand>
    <CHeaderNav class="d-md-down-none mr-auto">
      <CHeaderNavItem class="px-3">
        <!-- <CBreadcrumb :items="items"/> -->
         <CBreadcrumbRouter class="border-0 mb-0"/>
        
        <!-- <CHeaderNavLink to="/dashboard">
         Dashboard
        </CHeaderNavLink> -->
      </CHeaderNavItem>
    </CHeaderNav>
    <CHeaderNav>
        <CHeaderNavItem>
          {{serve}}
        </CHeaderNavItem>
      <CHeaderNavItem class="px-3 c-d-legacy-none">
        <button
          @click="() => $store.commit('toggle', 'darkMode')"
          class="c-header-nav-btn"
        >
          <CIcon v-if="$store.state.darkMode" name="cil-sun"/>
          <CIcon v-else name="cil-moon"/>
        </button>

        
      </CHeaderNavItem>
      <TheHeaderDropdownTasks/>
      <TheHeaderDropdownMssgs/>
      <TheHeaderDropdownAccnt/>
    </CHeaderNav>

  </CHeader>
</template>

<script>
import TheHeaderDropdownAccnt from './TheHeaderDropdownAccnt'
// import TheHeaderDropdownNotif from './TheHeaderDropdownNotif'
// import TheHeaderDropdownTasks from './TheHeaderDropdownTasks'
// import TheHeaderDropdownMssgs from './TheHeaderDropdownMssgs'

export default {
  name: 'TheHeader',
  components: {
    TheHeaderDropdownAccnt,
    // TheHeaderDropdownNotif,
    // TheHeaderDropdownTasks,
    // TheHeaderDropdownMssgs
  },
   data: function () {
    return {
      serve:localStorage.getItem("Servelocation")
    };
  },
}
</script>
