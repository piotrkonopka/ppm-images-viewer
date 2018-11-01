#version 410

in vec3 fragmentColor;
in vec2 fragmentTexture;
uniform sampler2D texture0;

out vec4 color;

void main() {
	color  = texture(texture0, fragmentTexture) * vec4(fragmentColor, 1.0f);
}