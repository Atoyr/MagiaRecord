const MAGICAL_GIRLS = [
        {
          id: 1,
          name: 'tamaki iroha',
          attribute: 'light',
          type: 'heel',
          hp: 2400,
          attack: 400,
          deffence: 300,
        },
        {
          id: 2,
          name: 'nanami yachiyo',
          attribute: 'watter',
          type: 'balance',
          hp: 2500,
          attack: 400,
          deffence: 300,
        },
        {
          id: 3,
          name: 'yui tsuruno',
          attribute: 'fire',
          type: 'magia',
          hp: 2400,
          attack: 400,
          deffence: 300,
        },
        {
          id: 4,
          name: 'mitsuki ferishia',
          attribute: 'dark',
          type: 'attack',
          hp: 2400,
          attack: 400,
          deffence: 300,
        },
        {
          id: 5,
          name: 'futaba sana',
          attribute: 'folest',
          type: 'deffence',
          hp: 2400,
          attack: 400,
          deffence: 400,
        },
      ]

export const state = () => ({
  magicalGirls : []
})

export const mutations = {
  clearMagicalGirls(state) {
    state.magicalGirls = []
  },
  updateMagicalGirls(state, magicalGirls) {
    state.magicalGirls = magicalGirls
  }
}

export const actions = {
  fetchMagicalGirls({commit}) {
    commit('updateMagicalGirls', MAGICAL_GIRLS)
  }
}
