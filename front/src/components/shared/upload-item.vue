<template>
  <tr
    class="upload-item"
    :data-file-id="upload.id"
    :key="upload.id"
  >
            <td class="file-name">
              {{ upload.name || '' }}              
            </td>
            <td class="size">{{ upload.size | formatSize }}</td>
            <td class="type">{{ upload.type }}</td>
            <td class="progress">
              <span v-if="upload.error" class="error">
                {{upload.error | capitalize }} error
              </span>
              <span v-else-if="upload.success || upload.active">
                <progress max='100' v-bind:value='progressStatus'></progress>
              </span>
              <span v-else>
                Pending
              </span>
            </td>
  </tr>
</template>

<script>
import { formatSize, capitalize } from '../../utils'

export default {
  props: ['item'],
  name: 'upload-item',
  filters: { formatSize, capitalize },

  data () {
    return {
      parentSongList: null
    }
  },

  computed: {
    /**
     * A shortcut to access the current vm's song (instead of this.item.song).
     * @return {Object}
     */
     fsize () {
      return this.item.upload.size
     },

     progressStatus: function () {
      return Math.trunc(this.item.upload.progress)
     },

    upload () {
      return this.item.upload
    }

  },

  methods: {

  }
}
</script>

<style lang="scss">
@import "../../assets/sass/partials/_vars.scss";
@import "../../assets/sass/partials/_mixins.scss";

.upload-item {
  border-bottom: 1px solid $color2ndBgr;
  height: 35px;

  html.no-touchevents &:hover {
    background: rgba(255, 255, 255, .05);
  }

  .type, .track-number {
    color: $color2ndText;
  }

  .file-name a e {
    min-width: 192px;
  }

  .play {
    max-width: 32px;
    opacity: .5;

    i {
      font-size: 1.5rem;
    }
  }

.progress {
  .error {
    color: red
  }
}

progress {  
    display:  block;
    //margin:  0.5em auto;
    background:  transparent;
    border:  0px solid #a0a0a0;
    height:  20px;
    width:  100%;
    position:  relative;
    &:after {
      color:  white;
      content:  attr(value);
      display:  block;
      position:  absolute;
      left:  0.25em;
      top:  0;
      height:  1.5rem;
      line-height:  1.5rem;
    }
}  
progress::-webkit-progress-bar {  
    background:  transparent;
    border:  0px solid #a0a0a0;
}  
progress::-webkit-progress-value {  
background:  #56a052;
}  
progress::-moz-progress-bar {  
background:  transparent;
}  


}

</style>
