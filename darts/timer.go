components {
  id: "timer"
  component: "/darts/timer.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"hg_100\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/balloons.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite-glow"
  type: "sprite"
  data: "default_animation: \"hg_glow\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/balloons.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.01
  }
}
