  <template>
  <div class="c-app" :class="{ 'c-dark-theme': $store.state.darkMode }"  > 
  
     <div>
       <!-- Check Connection Status -->
        <CToaster position="top-center" >
        <CToast :show="ConnectivityAlertMessage" color="danger"  align="center">
          <div class="d-flex text-center" align="center" style="width:60%; height:50%;">
          Connection failure : Offline
          </div>
        </CToast>
      </CToaster>
      <!-- LongTime No Repose user || Session Auto logout-->
       <CToaster position="top-center">
      <CToast :show="alertboxSuccess" color="info" align="center">
        <div class="d-flex text-center" >
            Your session will expire in <p class="p-0 m-0" style="color:white; padding-left:10px;" id="timeOut">&nbsp; 0 </p> ! 
           <div class="spinner-grow" role="status">
           <span class="sr-only" size="sm"> Loading...</span>
        </div>
      <CButton color="yahoo" @click="resetTimer">Clear</CButton>
        </div>
      </CToast>
    </CToaster>
    </div>
<!-- ++++++++++++ Master Section+++++++++++++ -->
    <TheSidebar/>
    <CWrapper>
      <TheHeader/>
      <div class="c-body" style="height:100%">
        <main class="c-main" style="height:100%;padding:0;">
          <CContainer fluid  style="height:100%">
            <transition name="fade" mode="out-in">
              <router-view :key="$route.path"></router-view>
            </transition>
          </CContainer>
        </main>
      </div>
      <TheFooter/>
    </CWrapper>
<!-- ++++++++++++Master Section+++++++++++++++ -->
  </div>
</template>

<script>
import TheSidebar from './TheSidebar'
import TheHeader from './TheHeader'
import TheFooter from './TheFooter'
import offline from 'v-offline';

export default {
  name: 'TheContainer',
  inject: ['api'],
 
  components: {
    TheSidebar,
    TheHeader,
    TheFooter,
    offline
  },

  data: function(){
    return{
    alertboxSuccess:false,
    events:['click','mousemove','mousedown','scroll','load','keypress',],
    warningTimer:null,
    logoutTimer:null,
    ConnectivityAlertMessage:null,
  }
  }, 

  created() {
    this.startup()
  },

 mounted(){
  this.events.forEach( function(event){
      window.addEventListener(event,this.test)
    },this)
    this.setTimer();
    // For detect connection status
    this.events.forEach( function(event){
      window.addEventListener(event,this.connectivity)
    },this)
  },

  destroyed(){
     this.events.forEach( function(event){
      window.removeEventListener(event,this.resetTimer)
    },this)
    this.resetTimer();

  },
  methods: {
    async startup () {
      try {
        var response = await this.api.get('/username')
        this.username = response.data
        if (this.username === null) this.$router.push('/login')
      } catch (error) {
        if (error.response.status === 401) {
          this.$router.push('/login')
        }
      }
    },
    async logout () {
      await this.api.get('/logout')
      this.$router.push('/login')
    },
   
// ==== Detect to Connectivity Status ====
  connectivity:function(){
    if(navigator.onLine === false){
      setTimeout(this.AlertMessage,3*1000);
    } else if(navigator.onLine === true){
      this.ConnectivityAlertMessage = false;
    }
  }, 

  AlertMessage: function(){
    this.ConnectivityAlertMessage = true;
  },
//========================================
  test:function(e){
    // console.log(e)
  },
  
  // 
  // SESSION LOG OUT
  //  timer(){
  //      var IdealTimeOut = 10; //10 seconds
  //       var idleSecondsTimer = null;
  //       var idleSecondsCounter = 0;
  //       document.onclick = function () { idleSecondsCounter = 0; };
  //       document.onmousemove = function () { idleSecondsCounter = 0; };
  //       document.onkeypress = function () { idleSecondsCounter = 0; };
  //       idleSecondsTimer = window.setInterval(CheckIdleTime, 1000);

  //       function CheckIdleTime() {
  //           idleSecondsCounter++;
  //           var oPanel = document.getElementById("timeOut");
  //           if (oPanel) {
  //               oPanel.innerHTML = (IdealTimeOut - idleSecondsCounter);
  //           }
  //           if (idleSecondsCounter >= IdealTimeOut) {
  //               window.clearInterval(idleSecondsTimer);
  //                 // await this.api.get('/logout')
  //               // this.$router.push('/login')
  //               setImmediate(sessionout)
  //               function sessionout(){
  //                 // alert("JOOJO")
  //                 // this.alertboxSuccess=false
  //                 // this.logout()
  //                 // this.$router.push('/login')
  //                 // window.location('/login')
  //               }
                
  //           }
  //       }
  //  },
  //  SET Timer call Warning Message
   setTimer: function(){
    //  this.warningTimer = setTimeout(this.warningMessage, 4 * 1000)
    //  this.logoutTimer = setTimeout(this.logout, 11 * 1000)
     this.alertboxSuccess=false
   },
   warningMessage:function(){
    // this.alertboxSuccess=true
    alert("session logout")
    // this.timer()
    // document.getElementById('timeOut').submit();
   },

   resetTimer:function(){
     this.clearTimeout(this.warningTimer)
     this.clearTimeout(this.logoutTime)
     this.setTimer()
    // this.alertboxSuccess=false
   },

  //=========================== OFFLINE ==================


  // offline:function(){
  //   window.addEventListener('load', function() { 
  //   var status = document.getElementById("status");
  //   var log = document.getElementById("log");

  // function updateOnlineStatus(event) {
  //   var condition = navigator.onLine ? "online" : "offline";

  //   status.className = condition;
  //   status.innerHTML = condition.toUpperCase();

  //   log.insertAdjacentHTML("beforeend", "Event: " + event.type + "; Status: " + condition);
  // }

  // window.addEventListener('online',  updateOnlineStatus);
  // window.addEventListener('offline', updateOnlineStatus);
  // });
  // },
  onFunction(){
    alert("online")
  }

  }
}
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>