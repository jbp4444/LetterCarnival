embedded_components {
  id: "background"
  type: "sprite"
  data: "default_animation: \"circle256half2\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/horses.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "letter1"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "color {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "}\n"
  "text: \"A\"\n"
  "font: \"/shooting/labelfont.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    y: 80.0
    z: 0.05
  }
}
embedded_components {
  id: "letter2"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "color {\n"
  "  x: 0.0\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "}\n"
  "text: \"C\"\n"
  "font: \"/shooting/labelfont.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  position {
    y: -80.0
    z: 0.05
  }
  scale {
    x: -1.0
    y: -1.0
  }
}
embedded_components {
  id: "circle1"
  type: "sprite"
  data: "default_animation: \"circle32\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/horses.atlas\"\n"
  "}\n"
  ""
  position {
    y: 80.0
    z: 0.01
  }
}
embedded_components {
  id: "circle2"
  type: "sprite"
  data: "default_animation: \"circle32\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/horses.atlas\"\n"
  "}\n"
  ""
  position {
    y: -80.0
    z: 0.01
  }
}
