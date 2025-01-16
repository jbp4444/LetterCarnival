components {
  id: "mole"
  component: "/whack/mole.script"
}
embedded_components {
  id: "background"
  type: "sprite"
  data: "default_animation: \"Shape-Square-256\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/animals.atlas\"\n"
  "}\n"
  ""
  position {
    y: -280.0
  }
}
embedded_components {
  id: "label"
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
    y: -220.0
    z: 0.01
  }
  scale {
    x: 2.22
    y: 2.22
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"giraffe\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/animals.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.01
  }
}
embedded_components {
  id: "stick"
  type: "sprite"
  data: "default_animation: \"stick_metal\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/sticks_atlas.atlas\"\n"
  "}\n"
  ""
  position {
    y: -120.0
    z: -0.01
  }
  scale {
    x: 4.0
  }
}
