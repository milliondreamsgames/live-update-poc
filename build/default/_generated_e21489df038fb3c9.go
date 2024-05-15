components {
  id: "hi_res_art"
  component: "/hi_res/hi_res_art.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  property_decls {
  }
}
embedded_components {
  id: "art_fg_1"
  type: "sprite"
  data: "default_animation: \"pie_rat_fg\"\nmaterial: \"/builtins/materials/sprite.material\"\nblend_mode: BLEND_MODE_ALPHA\ntextures {\n  sampler: \"texture_sampler\"\n  texture: \"/hi_res/hi_res.atlas\"\n}\n"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
