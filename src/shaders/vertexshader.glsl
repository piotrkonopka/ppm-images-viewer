#version 410

layout(location = 0) in vec3 vertexPosition;
layout(location = 1) in vec3 vertexColor;
layout(location = 2) in vec2 vertexTexture;

out vec3 fragmentColor;
out vec2 fragmentTexture;

void main() {
	gl_Position = vec4(vertexPosition, 1.0);
    fragmentColor = vertexColor;
    fragmentTexture = vertexTexture;
}