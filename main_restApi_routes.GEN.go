package main

// WARNING: DO NOT edit this file, it would be overwritten by domain/0_generator_test.go

import (
	"coba/conf"
	"coba/domain"

	"go.opentelemetry.io/otel/trace"

	"github.com/gofiber/fiber/v2"
	//"github.com/kokizzu/gotro/S"
)

func webApiInitRoutes(app *fiber.App) *domain.Domain {
	var (
		vdomain = domain.NewDomain()
	)

	app.All(conf.API_PREFIX+domain.Health_Url, func(ctx *fiber.Ctx) error {
		url := domain.Health_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.Health_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.Health(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.StoreCartItemsAdd_Url, func(ctx *fiber.Ctx) error {
		url := domain.StoreCartItemsAdd_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.StoreCartItemsAdd_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.StoreCartItemsAdd(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.StoreInvoice_Url, func(ctx *fiber.Ctx) error {
		url := domain.StoreInvoice_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.StoreInvoice_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.StoreInvoice(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.StoreProducts_Url, func(ctx *fiber.Ctx) error {
		url := domain.StoreProducts_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.StoreProducts_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.StoreProducts(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.UserChangeEmail_Url, func(ctx *fiber.Ctx) error {
		url := domain.UserChangeEmail_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.UserChangeEmail_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.UserChangeEmail(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.UserChangePassword_Url, func(ctx *fiber.Ctx) error {
		url := domain.UserChangePassword_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.UserChangePassword_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.UserChangePassword(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.UserConfirmEmail_Url, func(ctx *fiber.Ctx) error {
		url := domain.UserConfirmEmail_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.UserConfirmEmail_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.UserConfirmEmail(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.UserForgotPassword_Url, func(ctx *fiber.Ctx) error {
		url := domain.UserForgotPassword_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.UserForgotPassword_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.UserForgotPassword(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.UserList_Url, func(ctx *fiber.Ctx) error {
		url := domain.UserList_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.UserList_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.UserList(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.UserLogin_Url, func(ctx *fiber.Ctx) error {
		url := domain.UserLogin_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.UserLogin_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.UserLogin(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.UserLogout_Url, func(ctx *fiber.Ctx) error {
		url := domain.UserLogout_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.UserLogout_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.UserLogout(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.UserProfile_Url, func(ctx *fiber.Ctx) error {
		url := domain.UserProfile_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.UserProfile_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.UserProfile(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.UserRegister_Url, func(ctx *fiber.Ctx) error {
		url := domain.UserRegister_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.UserRegister_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.UserRegister(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.UserResetPassword_Url, func(ctx *fiber.Ctx) error {
		url := domain.UserResetPassword_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.UserResetPassword_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.UserResetPassword(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	app.All(conf.API_PREFIX+domain.XXX_Url, func(ctx *fiber.Ctx) error {
		url := domain.XXX_Url
		tracerCtx, span := conf.T.Start(ctx.Context(), url, trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		in := domain.XXX_In{}
		if err := webApiParseInput(ctx, &in.RequestCommon, &in, url); err != nil {
			return err
		}
		in.FromFiberCtx(ctx, tracerCtx)
		out := vdomain.XXX(&in)
		out.ToFiberCtx(ctx, &in.RequestCommon, &in)
		return in.ToFiberCtx(ctx, out)
	})

	return vdomain
}

// WARNING: DO NOT edit this file, it would be overwritten by domain/0_generator_test.go
