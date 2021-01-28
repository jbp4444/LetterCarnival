components {
  id: "cloud1"
  component: "/obstacles/cloud.script"
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
embedded_components {
  id: "stick"
  type: "sprite"
  data: "tile_set: \"/assets/sticks_atlas.atlas\"\n"
  "default_animation: \"stick_metal_outline\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 64.0
    z: -0.05
  }
  rotation {
    x: 1.0
    y: 0.0
    z: 0.0
    w: 6.123234E-17
  }
}
embedded_components {
  id: "cloud"
  type: "sprite"
  data: "tile_set: \"/assets/stall_atlas.atlas\"\n"
  "default_animation: \"cloud1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
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
