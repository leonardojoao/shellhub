<template>
  <v-app>
    <v-navigation-drawer
      v-if="isLoggedIn"
      v-model="drawer"
      app
      clipped
      dark
    >
      <v-list>
        <v-list-item
          v-for="item in items"
          :key="item.title"
          :to="item.path"
          two-line
        >
          <v-list-item-action>
            <v-icon v-text="item.icon" />
          </v-list-item-action>

          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar
      v-if="isLoggedIn"
      app
      clipped-left
      dark
      color="primary"
    >
      <v-app-bar-nav-icon
        class="hidden-lg-and-up"
        @click.stop="drawer = !drawer"
      />
      <router-link to="/">
        <v-img
          src="@/assets/logo-inverted.png"
          max-width="160"
        />
      </router-link>
      <span class="overline mt-3">BETA</span>
      <v-spacer />
      <v-menu
        transition="scale-transition"
        origin="top left"
      >
        <template v-slot:activator="{ on }">
          <v-chip v-on="on">
            <v-icon left>
              mdi-server
            </v-icon>
            My Device Fleet
            <v-icon right>
              mdi-chevron-down
            </v-icon>
          </v-chip>
        </template>
      </v-menu>
      <v-chip>
        <v-icon>help</v-icon>
      </v-chip>
      <v-chip>
        <v-icon>notifications</v-icon>
      </v-chip>
      <v-menu
        transition="scale-transition"
        origin="top right"
      >
        <template v-slot:activator="{ on }">
          <v-chip v-on="on">
            <v-icon left>
              mdi-account
            </v-icon>
            {{ $store.getters["auth/currentUser"] }}
            <v-icon right>
              mdi-chevron-down
            </v-icon>
          </v-chip>
        </template>

        <v-card>
          <v-list-item three-line>
            <v-list-item-content>
              <v-list-item-title class="mb-1">
                Tenant ID
              </v-list-item-title>
              <v-list-item-subtitle>
                <v-chip>
                  <span>
                    {{ tenant }}
                  </span>
                  <v-icon
                    v-clipboard="tenant"
                    v-clipboard:success="() => { copySnack = true; }"
                    right
                  >
                    mdi-content-copy
                  </v-icon>
                </v-chip>
              </v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>

          <v-card-actions>
            <v-btn
              small
              text
              @click="logout()"
            >
              Logout
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-menu>
    </v-app-bar>
    <v-main>
      <v-container
        class="pa-8"
        fluid
      >
        <router-view />
      </v-container>
      <v-snackbar
        v-model="copySnack"
        :timeout="3000"
      >
        Tenant ID copied to clipboard
      </v-snackbar>
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  data() {
    return {
      drawer: true,
      clipped: false,
      copySnack: false,
      items: [
        {
          icon: 'dashboard',
          title: 'Dashboard',
          path: '/',
        },
        {
          icon: 'devices',
          title: 'Devices',
          path: '/devices',
        },
        {
          icon: 'history',
          title: 'Sessions',
          path: '/sessions',
        },
      ],
    };
  },

  computed: {
    tenant() {
      return this.$store.getters['auth/tenant'];
    },

    isLoggedIn() {
      return this.$store.getters['auth/isLoggedIn'];
    },
  },

  methods: {
    logout() {
      this.$store.dispatch('auth/logout').then(() => {
        this.$router.push('/login');
      });
    },
  },
};

</script>

<style>
.v-list-active {
  border-left: 4px solid var(--v-primary-base);
}

.text-shadow {
  text-shadow: #000 0 0 6px;
  color: transparent;
}
</style>
