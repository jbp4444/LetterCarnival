components {
  id: "script"
  component: "/darts/balloon.script"
}
components {
  id: "sparkle"
  component: "/darts/sparkle.particlefx"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"balloon_green\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/balloons.atlas\"\n"
  "}\n"
  ""
  scale {
    y: -1.0
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
    y: -20.0
    z: 0.1
  }
  scale {
    x: 2.22
    y: 2.22
  }
}
embedded_components {
  id: "bg_circle"
  type: "sprite"
  data: "default_animation: \"circle128\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/balloons.atlas\"\n"
  "}\n"
  ""
  position {
    y: -20.0
    z: 0.01
  }
  scale {
    x: 0.7
    y: 0.7
  }
}
embedded_components {
  id: "sound"
  type: "sound"
  data: "sound: \"/assets/zapsplat/zapsplat_foley_balloon_pop_20568.ogg\"\n"
  ""
}
