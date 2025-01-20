components {
  id: "top_tgt"
  component: "/shooting/toptarget.script"
}
embedded_components {
  id: "front"
  type: "sprite"
  data: "default_animation: \"shipBeige_manned\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/aliens.atlas\"\n"
  "}\n"
  ""
  position {
    y: 17.0
  }
}
embedded_components {
  id: "stick"
  type: "sprite"
  data: "default_animation: \"stick_metal_outline\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/sticks_atlas.atlas\"\n"
  "}\n"
  ""
  position {
    x: -1.0
    y: 135.0
    z: -0.02
  }
  rotation {
    x: 1.0
    w: 6.123234E-17
  }
}
embedded_components {
  id: "back"
  type: "sprite"
  data: "default_animation: \"shipBeige_manned_back\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/aliens.atlas\"\n"
  "}\n"
  ""
  position {
    y: 17.0
  }
}
embedded_components {
  id: "circle"
  type: "sprite"
  data: "default_animation: \"circle32\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/ducks.atlas\"\n"
  "}\n"
  ""
  position {
    y: -19.0
    z: 0.025
  }
}
embedded_components {
  id: "label"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "text: \"Label\"\n"
  "font: \"/shooting/labelfont.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    y: -19.0
    z: 0.04
  }
}
embedded_components {
  id: "sound"
  type: "sound"
  data: "sound: \"/assets/kenney.nl/audio/impactGlass_heavy_000.ogg\"\n"
  ""
}
embedded_components {
  id: "movement"
  type: "sound"
  data: "sound: \"/assets/zapsplat/science_fiction_ufo_hover.ogg\"\n"
  ""
}
