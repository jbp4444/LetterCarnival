components {
  id: "duck_tgt"
  component: "/shooting/duck.script"
}
embedded_components {
  id: "front"
  type: "sprite"
  data: "default_animation: \"duck_yellow\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/ducks.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "stick"
  type: "sprite"
  data: "default_animation: \"stick_woodFixed_outline\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/sticks_atlas.atlas\"\n"
  "}\n"
  ""
  position {
    y: -108.0
    z: -0.02
  }
}
embedded_components {
  id: "back"
  type: "sprite"
  data: "default_animation: \"duck_back\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/ducks.atlas\"\n"
  "}\n"
  ""
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
    y: -18.0
    z: 0.02
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
    y: -18.0
    z: 0.04
  }
}
embedded_components {
  id: "sound"
  type: "sound"
  data: "sound: \"/assets/kenney.nl/audio/impactWood_heavy_000.ogg\"\n"
  ""
}
