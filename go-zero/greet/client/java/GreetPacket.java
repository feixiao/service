package com.xhb.logic.http.packet.greet;

import com.google.gson.Gson;
import com.xhb.commons.JSON;
import com.xhb.commons.JsonParser;
import com.xhb.core.network.HttpRequestClient;
import com.xhb.core.packet.HttpRequestPacket;
import com.xhb.core.response.HttpResponseData;
import com.xhb.logic.http.DeProguardable;

import org.jetbrains.annotations.NotNull;
import org.json.JSONObject;

public class GreetPacket extends HttpRequestPacket<GreetPacket.GreetPacketResponse> {

	private String name;

	public GreetPacket(String name, Request request) {
        super(request);
		this.request = request;
		this.name = name;
    }

	@Override
    public HttpRequestClient.Method requestMethod() {
        return HttpRequestClient.Method.GET;
    }

	@Override
    public String requestUri() {
        return "/greet/from/" + name;
    }

	@Override
    public GreetPacketResponse newInstanceFrom(JSON json) {
        return new GreetPacketResponse(json);
    }

	public static class GreetPacketResponse extends HttpResponseData {

		private Response responseData;

        GreetPacketResponse(@NotNull JSON json) {
            super(json);
            JSONObject jsonObject = json.asObject();
			if (JsonParser.hasKey(jsonObject, "data")) {
				Gson gson = new Gson();
				JSONObject dataJson = JsonParser.getJSONObject(jsonObject, "data");
				responseData = gson.fromJson(dataJson.toString(), Response.class);
			}
        }

		public Response getResponse () {
            return responseData;
        }
    }

	static class Request implements DeProguardable {
		private String name = "";

		@org.jetbrains.annotations.NotNull 
		public String getName() {
			return this.name;
		}

		public void setName(@org.jetbrains.annotations.NotNull String name) {
			this.name = name;
		}
	}

	static class Response implements DeProguardable {
		private int code;
		private String message = "";

		public int getCode() {
			return this.code;
		}

		public void setCode(int code) {
			this.code = code;
		}

		@org.jetbrains.annotations.NotNull 
		public String getMessage() {
			return this.message;
		}

		public void setMessage(@org.jetbrains.annotations.NotNull String message) {
			this.message = message;
		}
	}
}
