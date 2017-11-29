import config from '../config'
import artist from './artist'

export default {
  artist,
  id: "",
  artist_id: "",
  name: '',
  cover: config.unknownCover,
  playCount: 0,
  length: 0,
  fmtLength: '00:00',
  songs: []
}
