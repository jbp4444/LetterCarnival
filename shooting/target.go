components {
  id: "duck_tgt"
  component: "/shooting/target.script"
}
embedded_components {
  id: "front"
  type: "sprite"
  data: "default_animation: \"target_colored\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/ducks.atlas\"\n"
  "}\n"
  ""
  position {
    y: 17.0
  }
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
    x: -1.0
    y: -108.0
    z: -0.02
  }
}
embedded_components {
  id: "back"
  type: "sprite"
  data: "default_animation: \"target_back\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/ducks.atlas\"\n"
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
    y: 17.0
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
    y: 16.0
    z: 0.04
  }
}
embedded_components {
  id: "sound"
  type: "sound"
  data: "sound: \"/assets/kenney.nl/audio/impactMetal_heavy_000.ogg\"\n"
  ""
}
