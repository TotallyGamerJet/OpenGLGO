package main

import "github.com/go-gl/mathgl/mgl32"

const(
	VERTEX_FILE = `#version 450 core

in vec3 position;
in vec2 textureCoords;
in vec3 normal;

out vec2 pass_textureCoords;
out vec3 surfaceNormal;
out vec3 toLightVector;
out vec3 toCameraVector;
out float visibility;

uniform mat4 transformationMatrix;
uniform mat4 projectionMatrix;
uniform mat4 viewMatrix;
uniform vec3 lightPosition;

uniform float useFakeLighting;

const float density = 0.0035;
const float gradient = 5.0;

void main(void) {

	vec4 worldPosition = transformationMatrix * vec4(position, 1.0);
	vec4 positionRelativeToCam = viewMatrix * worldPosition;
	gl_Position = projectionMatrix * positionRelativeToCam;
	pass_textureCoords = textureCoords;

	vec3 actualNormal = normal;
	if(useFakeLighting > 0.5) {
		actualNormal = vec3(0.0,1.0,0.0);
	}

	surfaceNormal = (transformationMatrix * vec4(actualNormal, 0.0)).xyz;
	toLightVector =  lightPosition - worldPosition.xyz;
	toCameraVector = (inverse(viewMatrix) * vec4(0.0,0.0,0.0,1.0)).xyz - worldPosition.xyz;

	float distance = length(positionRelativeToCam.xyz);
	visibility = exp(-pow((distance*density),gradient));
	visibility = clamp(visibility,0.0,1.0);

}` + "\x00"
	FRAGMENT_FILE = `#version 450 core

in vec2 pass_textureCoords;
in vec3 surfaceNormal;
in vec3 toLightVector;
in vec3 toCameraVector;
in float visibility;

out vec4 out_Color;

uniform sampler2D textureSampler;
uniform vec3 lightColour;
uniform float shineDamper;
uniform float reflectivity;
uniform vec3 skyColour;

void main(void) {

	vec3 unitNormal = normalize(surfaceNormal);
	vec3 unitLightVector = normalize(toLightVector);

	float nDot1 = dot(unitNormal,unitLightVector);
	float brightness = max(nDot1,0.2);
	vec3 diffuse = brightness * lightColour;

	vec3 unitVectorToCamera = normalize(toCameraVector);
	vec3 lightDirection = -unitLightVector;
	vec3 reflectedLightDirection = reflect(lightDirection,unitNormal);

	float specularFactor = dot(reflectedLightDirection, unitVectorToCamera);
	specularFactor = max(specularFactor,0.0);
	float dampedFactor = pow(specularFactor,shineDamper);
	vec3 finalSpecular = dampedFactor * reflectivity * lightColour;

	vec4 textureColour = texture(textureSampler, pass_textureCoords);
	if(textureColour.a<0.5) {
		discard;
	}

    out_Color = vec4(diffuse, 1.0) * textureColour + vec4(finalSpecular, 1.0);
	out_Color = mix(vec4(skyColour,1.0),out_Color,visibility);

}` + "\x00"
)

type StaticShader struct {
	ShaderProgram
}

func newStaticShader() StaticShader {
	return StaticShader{newShaderProgram(VERTEX_FILE, FRAGMENT_FILE, bindAttributes, getAllUniformLocations)}
}

func (s StaticShader) loadTransformationMatrix(matrix mgl32.Mat4) {
	s.loadMatrix(s.uniformLocations["transformationMatrix"], matrix)
}

func (s StaticShader) loadProjectionMatrix(matrix mgl32.Mat4) {
	s.loadMatrix(s.uniformLocations["projectionMatrix"], matrix)
}

func (s StaticShader) loadViewMatrix(camera Camera) {
	viewMatrix := createViewMatrix(camera)
	s.loadMatrix(s.uniformLocations["viewMatrix"], viewMatrix)
}

func (s StaticShader) loadLight(light Light) {
	s.loadVector(s.uniformLocations["lightPosition"], light.position)
	s.loadVector(s.uniformLocations["lightColour"], light.colour)
}

func (s StaticShader) loadShineVariables(damper, reflectivity float32) {
	s.loadFloat(s.uniformLocations["shineDamper"], damper)
	s.loadFloat(s.uniformLocations["reflectivity"], reflectivity)
}

func (s StaticShader) loadFakeLighting(fakeLighting bool) {
	s.loadBoolean(s.uniformLocations["useFakeLighting"], fakeLighting)
}

func (s StaticShader) loadSkyColour(r,g,b float32) {
	s.loadVector(s.uniformLocations["skyColour"], mgl32.Vec3{r,g,b})
}

func bindAttributes(program ShaderProgram) {
	program.bindAttribute(0, "position")
	program.bindAttribute(1, "textureCoords")
	program.bindAttribute(2, "normal")

}

func getAllUniformLocations(program ShaderProgram) {
	program.uniformLocations["transformationMatrix"] = program.getUniformLocation("transformationMatrix")
	program.uniformLocations["projectionMatrix"] = program.getUniformLocation("projectionMatrix")
	program.uniformLocations["viewMatrix"] = program.getUniformLocation("viewMatrix")
	program.uniformLocations["lightPosition"] = program.getUniformLocation("lightPosition")
	program.uniformLocations["lightColour"] = program.getUniformLocation("lightColour")
	program.uniformLocations["shineDamper"] = program.getUniformLocation("shineDamper")
	program.uniformLocations["reflectivity"] = program.getUniformLocation("reflectivity")
	program.uniformLocations["useFakeLighting"] = program.getUniformLocation("useFakeLighting")
	program.uniformLocations["skyColour"] = program.getUniformLocation("skyColour")

}