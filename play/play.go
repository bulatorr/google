package play

import (
   "154.pages.dev/protobuf"
   "bytes"
   "compress/gzip"
   "encoding/base64"
   "fmt"
   "net/http"
   "time"
)

// com.roku.web.trc
const Leanback = "android.software.leanback"

const android_api = 31

// the device actually uses 0x30000, but some apps require a higher version:
// com.axis.drawingdesk.v3
// so lets lie for now
const gl_es_version = 0x30002

const google_play_store = 82941300

// developer.android.com/ndk/guides/abis
var ABI = []string{
   "arm64-v8a",
   "armeabi-v7a",
   "x86",
}

var Device = GoogleDevice{
   Feature: []string{
      "android.hardware.sensor.proximity", "android.hardware.telephony.ims.singlereg", "android.hardware.sensor.accelerometer", "android.software.controls", "android.hardware.faketouch", "android.software.telecom", "android.hardware.telephony.subscription", "android.hardware.telephony.euicc", "android.hardware.usb.accessory", "android.hardware.telephony.data", "android.hardware.sensor.dynamic.head_tracker", "android.software.backup", "android.hardware.touchscreen", "android.hardware.touchscreen.multitouch", "android.software.erofs", "android.software.print", "android.software.activities_on_secondary_displays", "android.hardware.wifi.rtt", "com.google.android.feature.PIXEL_2017_EXPERIENCE", "android.software.voice_recognizers", "android.software.picture_in_picture", "android.hardware.fingerprint", "android.hardware.sensor.gyroscope", "android.hardware.audio.low_latency", "android.software.vulkan.deqp.level", "android.software.cant_save_state", "com.google.android.feature.PIXEL_2018_EXPERIENCE", "android.hardware.security.model.compatible", "android.hardware.telephony.messaging", "com.google.android.feature.PIXEL_2019_EXPERIENCE", "android.hardware.telephony.calling", "android.hardware.opengles.aep", "org.lineageos.livedisplay", "android.hardware.bluetooth", "android.software.window_magnification", "android.hardware.telephony.radio.access", "android.hardware.camera.autofocus", "android.hardware.telephony.gsm", "android.hardware.telephony.ims", "android.software.incremental_delivery", "android.hardware.se.omapi.ese", "android.software.opengles.deqp.level", "vendor.android.hardware.camera.preview-dis.front", "com.google.android.feature.PIXEL_2022_MIDYEAR_EXPERIENCE", "android.hardware.camera.concurrent", "android.hardware.usb.host", "android.hardware.audio.output", "android.software.verified_boot", "android.hardware.camera.flash", "android.hardware.camera.front", "android.hardware.sensor.hifi_sensors", "android.hardware.se.omapi.uicc", "android.hardware.strongbox_keystore", "android.hardware.screen.portrait", "android.hardware.nfc", "com.nxp.mifare", "com.google.android.feature.PIXEL_2021_MIDYEAR_EXPERIENCE", "android.hardware.sensor.stepdetector", "android.software.home_screen", "android.hardware.context_hub", "vendor.android.hardware.camera.preview-dis.back", "android.hardware.microphone", "android.software.autofill", "org.lineageos.hardware", "org.lineageos.globalactions", "android.software.securely_removes_users", "com.google.android.feature.PIXEL_EXPERIENCE", "android.hardware.bluetooth_le", "android.hardware.sensor.compass", "android.hardware.touchscreen.multitouch.jazzhand", "android.hardware.sensor.barometer", "android.software.app_widgets", "com.google.android.feature.PIXEL_2020_MIDYEAR_EXPERIENCE", "android.hardware.telephony.carrierlock", "android.software.input_methods", "android.hardware.sensor.light", "android.hardware.vulkan.version", "android.software.companion_device_setup", "android.software.device_admin", "android.hardware.wifi.passpoint", "android.hardware.camera", "org.lineageos.trust", "android.hardware.device_unique_attestation", "android.hardware.screen.landscape", "android.software.device_id_attestation", "android.hardware.ram.normal", "com.google.android.feature.PIXEL_2019_MIDYEAR_EXPERIENCE", "android.software.managed_users", "android.software.webview", "android.hardware.sensor.stepcounter", "android.hardware.camera.capability.manual_post_processing", "android.hardware.camera.any", "android.hardware.camera.capability.raw", "android.hardware.vulkan.compute", "android.hardware.touchscreen.multitouch.distinct", "android.hardware.location.network", "android.software.cts", "android.hardware.camera.capability.manual_sensor", "android.software.app_enumeration", "android.hardware.camera.level.full", "android.hardware.identity_credential", "android.hardware.wifi.direct", "android.software.live_wallpaper", "com.google.android.feature.GOOGLE_EXPERIENCE", "android.software.ipsec_tunnels", "org.lineageos.settings", "android.hardware.audio.pro", "android.hardware.nfc.hcef", "android.hardware.location.gps", "android.software.midi", "android.hardware.nfc.any", "android.hardware.nfc.ese", "android.hardware.nfc.hce", "android.hardware.hardware_keystore", "com.google.android.feature.PIXEL_2020_EXPERIENCE", "android.hardware.telephony.euicc.mep", "android.hardware.wifi", "android.hardware.location", "android.hardware.vulkan.level", "com.google.android.feature.PIXEL_2021_EXPERIENCE", "android.hardware.keystore.app_attest_key", "android.hardware.wifi.aware", "com.google.android.feature.PIXEL_2022_EXPERIENCE", "android.software.secure_lock_screen", "android.hardware.telephony", "android.software.file_based_encryption",
   },
   Library: []string{
      "android.test.base", "android.test.mock", "android.hidl.manager-V1.0-java", "google-ril", "libedgetpu_client.google.so", "libedgetpu_util.so", "android.hidl.base-V1.0-java", "com.google.android.camera.experimental2022", "libOpenCL-pixel.so", "com.android.location.provider", "oemrilhook", "android.net.ipsec.ike", "com.android.future.usb.accessory", "android.ext.shared", "javax.obex", "com.google.android.gms", "lib_aion_buffer.so", "libgxp.so", "gxp_metrics_logger.so", "android.test.runner", "org.apache.http.legacy", "com.android.cts.ctsshim.shared_library", "com.android.nfc_extras", "com.android.media.remotedisplay", "com.android.mediadrm.signer", "android.system.virtualmachine",
   },
   Texture: []string{
      "GL_ANDROID_extension_pack_es31a", "GL_ARM_mali_program_binary", "GL_ARM_mali_shader_binary", "GL_ARM_rgba8", "GL_ARM_shader_framebuffer_fetch", "GL_ARM_shader_framebuffer_fetch_depth_stencil", "GL_ARM_texture_unnormalized_coordinates", "GL_EXT_EGL_image_array", "GL_EXT_YUV_target", "GL_EXT_blend_minmax", "GL_EXT_buffer_storage", "GL_EXT_clip_control", "GL_EXT_color_buffer_float", "GL_EXT_color_buffer_half_float", "GL_EXT_copy_image", "GL_EXT_debug_marker", "GL_EXT_discard_framebuffer", "GL_EXT_disjoint_timer_query", "GL_EXT_draw_buffers_indexed", "GL_EXT_draw_elements_base_vertex", "GL_EXT_external_buffer", "GL_EXT_float_blend", "GL_EXT_geometry_shader", "GL_EXT_gpu_shader5", "GL_EXT_multisampled_render_to_texture", "GL_EXT_multisampled_render_to_texture2", "GL_EXT_occlusion_query_boolean", "GL_EXT_primitive_bounding_box", "GL_EXT_protected_textures", "GL_EXT_read_format_bgra", "GL_EXT_robustness", "GL_EXT_sRGB", "GL_EXT_sRGB_write_control", "GL_EXT_shader_framebuffer_fetch", "GL_EXT_shader_io_blocks", "GL_EXT_shader_non_constant_global_initializers", "GL_EXT_shader_pixel_local_storage", "GL_EXT_shadow_samplers", "GL_EXT_tessellation_shader", "GL_EXT_texture_border_clamp", "GL_EXT_texture_buffer", "GL_EXT_texture_compression_astc_decode_mode", "GL_EXT_texture_compression_astc_decode_mode_rgb9e5", "GL_EXT_texture_cube_map_array", "GL_EXT_texture_filter_anisotropic", "GL_EXT_texture_format_BGRA8888", "GL_EXT_texture_rg", "GL_EXT_texture_sRGB_R8", "GL_EXT_texture_sRGB_RG8", "GL_EXT_texture_sRGB_decode", "GL_EXT_texture_storage", "GL_EXT_texture_type_2_10_10_10_REV", "GL_EXT_unpack_subimage", "GL_KHR_blend_equation_advanced", "GL_KHR_blend_equation_advanced_coherent", "GL_KHR_debug", "GL_KHR_robust_buffer_access_behavior", "GL_KHR_robustness", "GL_KHR_texture_compression_astc_hdr", "GL_KHR_texture_compression_astc_ldr", "GL_KHR_texture_compression_astc_sliced_3d", "GL_OES_EGL_image", "GL_OES_EGL_image_external", "GL_OES_EGL_image_external_essl3", "GL_OES_EGL_sync", "GL_OES_blend_equation_separate", "GL_OES_blend_func_separate", "GL_OES_blend_subtract", "GL_OES_byte_coordinates", "GL_OES_compressed_ETC1_RGB8_texture", "GL_OES_compressed_paletted_texture", "GL_OES_copy_image", "GL_OES_depth24", "GL_OES_depth_texture", "GL_OES_depth_texture_cube_map", "GL_OES_draw_buffers_indexed", "GL_OES_draw_elements_base_vertex", "GL_OES_draw_texture", "GL_OES_element_index_uint", "GL_OES_extended_matrix_palette", "GL_OES_fbo_render_mipmap", "GL_OES_fixed_point", "GL_OES_framebuffer_object", "GL_OES_geometry_shader", "GL_OES_get_program_binary", "GL_OES_gpu_shader5", "GL_OES_mapbuffer", "GL_OES_matrix_get", "GL_OES_matrix_palette", "GL_OES_packed_depth_stencil", "GL_OES_point_size_array", "GL_OES_point_sprite", "GL_OES_primitive_bounding_box", "GL_OES_query_matrix", "GL_OES_read_format", "GL_OES_required_internalformat", "GL_OES_rgb8_rgba8", "GL_OES_sample_shading", "GL_OES_sample_variables", "GL_OES_shader_image_atomic", "GL_OES_shader_io_blocks", "GL_OES_shader_multisample_interpolation", "GL_OES_single_precision", "GL_OES_standard_derivatives", "GL_OES_stencil8", "GL_OES_stencil_wrap", "GL_OES_surfaceless_context", "GL_OES_tessellation_shader", "GL_OES_texture_3D", "GL_OES_texture_border_clamp", "GL_OES_texture_buffer", "GL_OES_texture_compression_astc", "GL_OES_texture_cube_map", "GL_OES_texture_cube_map_array", "GL_OES_texture_float_linear", "GL_OES_texture_mirrored_repeat", "GL_OES_texture_npot", "GL_OES_texture_stencil8", "GL_OES_texture_storage_multisample_2d_array", "GL_OES_vertex_array_object", "GL_OES_vertex_half_float", "GL_OVR_multiview", "GL_OVR_multiview2", "GL_OVR_multiview_multisampled_render_to_texture",
   },
}

func compress_gzip(p []byte) ([]byte, error) {
   var b bytes.Buffer
   w := gzip.NewWriter(&b)
   if _, err := w.Write(p); err != nil {
      return nil, err
   }
   if err := w.Close(); err != nil {
      return nil, err
   }
   return b.Bytes(), nil
}

func user_agent(req *http.Request, single bool) {
   var b []byte
   // `sdk` is needed for `/fdfe/delivery`
   b = append(b, "Android-Finsky (sdk="...)
   // with `/fdfe/acquire`, requests will be rejected with certain apps, if the
   // device was created with too low a version here:
   b = fmt.Append(b, android_api)
   b = append(b, ",versionCode="...)
   // for multiple APKs just tell the truth. for single APK we have to lie.
   // below value is the last version that works.
   if single {
      b = fmt.Append(b, 80919999)
   } else {
      b = fmt.Append(b, google_play_store)
   }
   b = append(b, ')')
   req.Header.Set("User-Agent", string(b))
}

func (g GoogleAuth) authorization(req *http.Request) {
   req.Header.Set("authorization", "Bearer " + g.get_auth())
}

func (g GoogleCheckin) x_dfe_device_id(req *http.Request) error {
   id, err := g.device_id()
   if err != nil {
      return err
   }
   req.Header.Set("x-dfe-device-id", fmt.Sprintf("%x", id))
   return nil
}

func (g GoogleCheckin) x_ps_rh(req *http.Request) error {
   id, err := g.device_id()
   if err != nil {
      return err
   }
   var m protobuf.Message
   m.Add(1, func(m *protobuf.Message) {
      m.Add(1, func(m *protobuf.Message) {
         m.Add(3, func(m *protobuf.Message) {
            m.AddBytes(1, fmt.Append(nil, id))
            m.Add(2, func(m *protobuf.Message) {
               now := time.Now().UnixMicro()
               m.AddBytes(1, fmt.Append(nil, now))
            })
         })
      })
   })
   data, err := compress_gzip(m.Encode())
   if err != nil {
      return err
   }
   req.Header.Set("x-ps-rh", base64.URLEncoding.EncodeToString(data))
   return nil
}

type GoogleDevice struct {
   ABI string
   Feature []string
   Library []string
   Texture []string
}

// play.google.com/store/apps/details?id=com.google.android.apps.youtube.unplugged
type StoreApp struct {
   ID string
   Version uint64
}

func (s StoreApp) APK(v string) string {
   var b []byte
   b = fmt.Append(b, s.ID, "-")
   if v != "" {
      b = fmt.Append(b, v, "-")
   }
   b = fmt.Append(b, s.Version, ".apk")
   return string(b)
}

func (s StoreApp) OBB(v uint64) string {
   var b []byte
   if v >= 1 {
      b = append(b, "patch"...)
   } else {
      b = append(b, "main"...)
   }
   b = fmt.Append(b, ".", s.Version, ".", s.ID, ".obb")
   return string(b)
}
