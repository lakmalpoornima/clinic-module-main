<template>
  <CDropdown
    inNav
    class="c-header-nav-items"
    placement="bottom-end"
    add-menu-classes="pt-0"
  >
    <template #toggler>
      <CHeaderNavLink>
        <div class="c-avatar">
          <img src="img/avatars/user.png" class="c-avatar-img" />
        </div>
      </CHeaderNavLink>
    </template>
    <CDropdownHeader tag="div" class="text-center" color="light">
      <strong>Settings</strong>
    </CDropdownHeader>
    <CDropdownItem> <CIcon name="cil-user" /> Profile </CDropdownItem>
    <CDropdownDivider />
    <CDropdownItem>
      <CIcon name="cil-shield-alt" /> Lock Account
    </CDropdownItem>
    <CDropdownItem @click="logout">
      <CIcon name="cil-lock-locked" @click="for_cancel()" /> Logout
    </CDropdownItem>
  </CDropdown>
</template>

<script>
export default {
  name: "TheHeaderDropdownAccnt",
  data() {
    return {
      itemsCount: 42,
    };
  },
  methods: {
    logout() {
      if (confirm("Are you sure to logout") == true) {
        window.location.href = "/";
        this.removeOldSession();
      } else {
        text = "You canceled!";
      }
    },
    /*===========================================================*/
    //    LOG OUT THEN FLUSH ALL THE STORE LOCAL STORAGE MEMORY
    /*===========================================================*/
    removeOldSession() {
      localStorage.clear();
    },
  },
};
</script>

<style scoped>
.c-icon {
  margin-right: 0.3rem;
}
</style>
