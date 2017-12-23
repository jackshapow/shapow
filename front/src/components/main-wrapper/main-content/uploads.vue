<template>
  <section id="uploadWrapper">

    <h1 class="heading">
      <span title="uploads">Uploads
        <span class="meta">
          <span v-show="transferStatus.uploadCount">
            {{ transferStatus.successCount }} of {{ transferStatus.uploadCount }} uploaded
          </span>
          <span v-show="transferStatus.remainCount">
            •
            {{ transferStatus.secondsLeftAvg | secondsToHis }} remaining
          </span>

          <span v-if="transferStatus.errorCount">
            •
            <span style="color:red">
            {{ transferStatus.errorCount | pluralize('error') }} 
            </span>
          </span>
        </span>
      </span>

      <div class="buttons">

        <file-upload
          class="btn btn-green btn-add"
          post-action="/api/songs"
          @input-file="inputFile"
          :multiple="true"
          :thread="5"
          :drop="true"
          :drop-directory="true"
          v-model="files"
          ref="upload">
          <i class="fa fa-plus"></i>
          ADD
        </file-upload>

      </div>
    </h1>

    <div v-show="$refs.upload && $refs.upload.dropActive" class="drop-active">
      <h3>Drop files to upload</h3>
    </div>

      <div v-if="uploads.length" class="upload-list-wrap main-scroll-wrap" :class="type"
        ref="wrapper"
        tabindex="1"
      >

        <table class="upload-list-header">
          <thead>
            <tr>
              <th class="file-name">Name
              </th>
              <th class="size">Size
              </th>
              <th class="type">Type
              </th>
              <th class="progress">Progress
              </th>
            </tr>
          </thead>
        </table>

        <virtual-scroller
          class="scroller"
          content-tag="table"
          :items="uploads"
          item-height="35"
          :renderers="renderers"
        />

      </div>

    </div>

    <div v-else class="none">
      <p>Nothing here yet. Drag and drop some files to add them to your collection.</p>
    </div>

  </section>

</template>

<script>
import { pluralize, formatSize, secondsToHis } from '../../../utils'
import { queueStore, songStore } from '../../../stores'
import FileUpload from 'vue-upload-component/src'
import uploadItem from '../../shared/upload-item.vue'

export default {
  components: { FileUpload, uploadItem },
  name: 'main-wrapper--main-content--queue',
  filters: { pluralize, formatSize, secondsToHis },
  props: ['type'],

  data () {
    return {
      renderers: Object.freeze({
        upload: uploadItem
      }),
      files: [],
      uploads: [],
      transferStatus: {secondsLeftAvg: 0, uploadCount: 0, errorCount: 0, remainCount: 0, activeCount: 0, successCount: 0},
    }
  },

  watch: {
    files () {
      this.generateUploads()
      this.updateTransferStatus()
    }
  },

  computed: {
  },

  methods: {

    updateTransferStatus () {
      // Grab latest data from files array
      // TODO: factor in error status codes
      var newStatus = this.files.reduce((a, b) => {
        if (!b.active && !b.success) { // File is pending transfer
          a.bytesLeft += b.size
        } else if (b.active && !b.success) { // File is being transferred
          a.bytesLeft += (((100 - b.progress) / 100) * b.size)
          a.bytesPerSecond += b.speed
        }

        // Set summary totals
        a.uploadCount += 1
        if (b.error) { a.errorCount += 1 }
        if (b.active) { a.activeCount += 1 }
        if (!b.success && !b.error) { a.remainCount += 1 }
        if (b.success) { a.successCount += 1 }

        return a
      }, {bytesLeft: 0, bytesPerSecond: 0, errorCount: 0, uploadCount: 0, remainCount: 0, activeCount: 0, successCount: 0, secondsLeftAvg: 0, secondsLeft: 0})

      // Stop calculating if no bytes received
      if (newStatus.bytesPerSecond == 0) {
        this.transferStatus = newStatus
        return true
      } 

      newStatus.secondsLeft = newStatus.bytesLeft / newStatus.bytesPerSecond    

      // If this is first status update then set a default previous status for smoothign function
      if (this.transferStatus.secondsLeftAvg == 0) {
        this.transferStatus.secondsLeftAvg = newStatus.secondsLeft
      }

      // Reset previous average if bandwidth is jumping around too much.
      newStatus.rangeDiff = (this.transferStatus.secondsLeftAvg-newStatus.secondsLeft)/newStatus.secondsLeft * 100
      if (Math.abs(newStatus.rangeDiff) > 100) {
        this.transferStatus.secondsLeftAvg = newStatus.secondsLeft
      }

      // Smooth out time updates
      var smoothing = 0.02
      newStatus.secondsLeftAvg = smoothing * newStatus.secondsLeft + (1-smoothing) * this.transferStatus.secondsLeftAvg;

      this.transferStatus = newStatus
      return true
    },

    generateUploads () {
      this.uploads = this.files.map(upload => {
        return {
          upload,
          type: 'upload'
        }
      })      
    },
    // Automatic file upload
    inputFile(newFile, oldFile) {
      if (Boolean(newFile) !== Boolean(oldFile) || oldFile.error !== newFile.error) {
        if (!this.$refs.upload.active) {
          this.$refs.upload.active = true
        }
      }
    }
  }
}
</script>

<style lang="scss">
@import "../../../assets/sass/partials/_vars.scss";
@import "../../../assets/sass/partials/_mixins.scss";

.drop-active {
  top: 0;
  bottom: 0;
  right: 0;
  left: 0;
  position: fixed;
  z-index: 9999;
  opacity: .6;
  text-align: center;
  background: #000;
}
.drop-active h3 {
  margin: -.5em 0 0;
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  -webkit-transform: translateY(-50%);
  -ms-transform: translateY(-50%);
  transform: translateY(-50%);
  font-size: 40px;
  color: #fff;
  padding: 0;
}

.file-uploads {
  overflow: hidden;
  position: relative;
  text-align: center;
  display: inline-block;
}
.file-uploads.file-uploads-html4 input[type="file"] {
  opacity: 0;
  font-size: 20em;
  z-index: 1;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  position: absolute;
  width: 100%;
  height: 100%;
}
.file-uploads.file-uploads-html5 input[type="file"] {
  overflow: hidden;
  position: fixed;
  width: 1px;
  height: 1px;
  z-index: -1;
  opacity: 0;
}


#uploadWrapper {
  .none {
    color: $color2ndText;
    padding: 16px 24px;

    a {
      color: $colorHighlight;
    }
  }

  label.btn {
    margin-bottom: 0;
    margin-right: 1rem;
  }

}


.upload-list-wrap {
  position: relative;
  padding: 8px 24px;

  .upload-list-header {
    position: absolute;
    top: 0;
    left: 24px;
    right: 24px;
    padding: 0 24px;
    background: #1b1b1b;
    z-index: 1;
    width: calc(100% - 48px);
  }

  table {
    width: 100%;
    table-layout: fixed;
  }

  tr.droppable {
    border-bottom-width: 3px;
    border-bottom-color: $colorGreen;
  }

  td, th {
    text-align: left;
    padding: 8px;
    vertical-align: middle;
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;

    &.file-name {
      width: 50%;
    }

    &.size {
      width: 100px;
      //text-align: right;
    }

    &.type {
      width: 100px;
    }

    &.progress {
      width: 30%;
    }

    &.play {
      display: none;

      html.touchevents & {
        display: block;
        position: absolute;
        top: 8px;
        right: 4px;
      }
    }
  }

  th {
    color: $color2ndText;
    letter-spacing: 1px;
    text-transform: uppercase;
    cursor: pointer;

    i {
      color: $colorHighlight;
      font-size: 1.2rem;
    }
  }

  /**
   * Since the Queue screen doesn't allow sorting, we reset the cursor style.
   */
  &.queue th {
    cursor: default;
  }

  .scroller {
    overflow: auto;
    position: absolute;
    top: 35px;
    left: 0;
    bottom: 0;
    right: 0;
    overflow-y: scroll;
    -webkit-overflow-scrolling: touch;

    .item-container {
      position: absolute;
      left: 24px;
      right: 24px;
    }

    .item {
      margin-bottom: 0;
    }
  }

  @media only screen and (max-width: 768px) {
    table, tbody, tr {
      display: block;
    }

    thead, tfoot {
      display: none;
    }

    .scroller {
      top: 0;
      bottom: 24px;

      .item-container {
        left: 12px;
        right: 12px;
      }
    }

    tr {
      padding: 8px 32px 8px 4px;
      position: relative;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
      color: $color2ndText;
      width: 100%;
    }

    td {
      display: inline;
      padding: 0;
      vertical-align: bottom;
      color: $colorMainText;

      &.album, &.time, &.track-number {
        display: none;
      }

      &.artist {
        color: $color2ndText;
        font-size: .9rem;
        padding: 0 4px;
      }
    }
  }
}
</style>
